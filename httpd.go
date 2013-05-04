package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type HtSession struct {
	sessionid string
	user      *User
}

type HtAuthorization struct {
	Username      string
	Authorization string
	Authtype      int // 0-Password,1-Session,2-apikey
	user 		 	User 
}

type HtMsg interface{}

type HtRequest struct {
	Auth 			HtAuthorization
	session       	*HtSession
	Message       	HtMsg
	Nada          	string
}

type HtMsgMsg struct {
	Message		string
    Servername 	string
    Nick 		string
    Recipient 	string
}

func initHttp(users *[]User) {

	authenticate := func(auth *HtAuthorization){
		// Authenticate()
		stub("Do proper authentication here")
		for _,v := range *users{
			fmt.Println("? ", auth.Username, "?=", v.username)
			if (v.username==auth.Username){
				fmt.Println("\n\n User match!")
				auth.user = v
				return
			}
		}

		fmt.Println("\n\n NO MATCH! \n\n")
	}

	// Send a message through /msg/
	// Example:
	/*
			{
				"Auth": {
			    	"Username": "dingolvrB"
			    }
			    ,
			   "Message" : {
			   		"Message": "LUVYa",
			        "Servername": "freenode",
			        "Nick": "dingolvrB1",
			        "Recipient": "#dingolove"
			    },"Nada":"whut"
			}	
	*/
	http.HandleFunc("/msg/", func(writer http.ResponseWriter, r *http.Request) {
		var req HtRequest
		var msg HtMsgMsg

		// FromRequest()
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &req)
		authenticate(&req.Auth)

		// FIXME: Find a better way to get to this...
		msgjson,_:=json.Marshal(req.Message)
		json.Unmarshal(msgjson, &msg)







		// Get Identity
		for _,v := range req.Auth.user.identities{
			fmt.Println("\nbeep:\n", v.Servername,"/",v.Nick ,"\n", msg.Servername,"/",msg.Nick,"\n")

			if (v.Servername == msg.Servername && v.Nick == msg.Nick) {
				// We found the identity.  Send the msg.
				v.connection.Privmsg(msg.Recipient, msg.Message)
			}
		}

		//req.user.identities[0].connection.Privmsg("#dingolove", msg.Message)

		jsn, _ := json.MarshalIndent(req, "", "      ")
		io.WriteString(writer, string(jsn))
		io.WriteString(writer, req.Auth.user.username)

	})

	http.HandleFunc("/history/", func(writer http.ResponseWriter, r *http.Request) {
		var req HtRequest
		var msg HtMsgMsg

		// FromRequest()
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &req)
		authenticate(&req.Auth)

		// FIXME: Find a better way to get to this...
		msgjson,_:=json.Marshal(req.Message)
		json.Unmarshal(msgjson, &msg)

		output,_ := json.MarshalIndent(&req.Auth.user.identities, "",  "      ")

		io.WriteString(writer, string(output))
	})

	http.HandleFunc("/sandbox/body", func(writer http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		io.WriteString(writer, fmt.Sprintf(r.Method, string(body)))
	})

	http.HandleFunc("/sandbox/json", func(writer http.ResponseWriter, r *http.Request) {
		var f interface{}
		body, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(body, &f)
		msg, err := json.MarshalIndent(&f, "", "    ")
		if err != nil {
			stub("RUHROH")
		}
		io.WriteString(writer, string(msg))

	})

	http.HandleFunc("/sandbox/jsonspecific", func(writer http.ResponseWriter, r *http.Request) {
		type Moar struct {
			What	string
			MoarThings	[]string
		}
		type Message struct {
			Name string
			Body string
			Time int64
			Things []int64
			Moar 	Moar

		}
		body, _ := ioutil.ReadAll(r.Body)
		var m Message
		json.Unmarshal(body, &m)

		io.WriteString(writer, fmt.Sprintf("",m.Moar.MoarThings))
	})

	http.Handle("/", http.FileServer(http.Dir("./pub")))

	http.ListenAndServe(":7776", nil)
}
