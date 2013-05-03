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
}

type HtMsg interface{}

type HtRequest struct {
	Authorization 	HtAuthorization
	session       	*HtSession
	user 			User 
	Message       	HtMsg
	Nada          	string
}

func initHttp(users *[]User) {

	//sessions := []*HtSession{}

	getreq := func(r *http.Request) HtRequest {
		body, _ := ioutil.ReadAll(r.Body)
		var req HtRequest
		json.Unmarshal(body, &req)

		stub("Do proper authentication here")

		for _,v := range *users{
			if (v.username==req.Authorization.Username){
				fmt.Println("\n\n User match!")
				req.user = v
			}
		}
		return req

	}

	http.HandleFunc("/msg/", func(writer http.ResponseWriter, r *http.Request) {
		req := getreq(r)

		req.user.identities[0].connection.Privmsg("#dingolove", req.Message.(string))




		jsn, _ := json.MarshalIndent(&req, "", "      ")
		io.WriteString(writer, string(jsn))
		io.WriteString(writer, req.user.username)

	})

	http.HandleFunc("/history/", func(writer http.ResponseWriter, r *http.Request) {
		io.WriteString(writer, "history!\n")
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

	http.ListenAndServe(":7776", nil)
}
