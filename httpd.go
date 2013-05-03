package main

import(
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type HtSession struct {
	sessionid		string
	user 			*User
}

type HtAuthObj struct {
	username 		string
	authorization	string
	authtype		int  // 0-Password,1-Session,2-apikey
}

type HtMsg interface {}

type HtReq struct {
	authorization	HtAuthObj
	session 		*HtSession
	message 		HtMsg
	nada			string
}

func initHttp(users *[]User){

	sessions := []*HtSession{}

	getreq := func( r *http.Request ) (HtReq,error){
 		body, err := ioutil.ReadAll(r.Body)
 		req := HtReq{}
 		json.Unmarshal(body, &req)

		if req.authorization.authtype == 0 {
			stub("Assume the password is right")

			stub("Create a session")


		} else if req.authorization.authtype == 1 {

			stub ("Search for matching session id")
			for _,v := range sessions{
				if (v.sessionid == req.authorization.authorization){
					req.session = v
				}
			}

		}

		return req, err
	}


    http.HandleFunc("/msg/", func (writer http.ResponseWriter, r *http.Request){
    	req , err := getreq(r)
    	if err!=nil{
    		io.WriteString(writer,"error. ")
	    	io.WriteString(writer,fmt.Sprintf("",err))
    	}
		io.WriteString(writer, req.nada)

	})

    http.HandleFunc("/history/", func (writer http.ResponseWriter, r *http.Request){
		io.WriteString(writer, "history!\n")
	})


    http.HandleFunc("/sandbox/body", func (writer http.ResponseWriter, r *http.Request){
 		body, _ := ioutil.ReadAll(r.Body)
    	io.WriteString(writer, fmt.Sprintf( r.Method, string(body) ))
	})

    http.HandleFunc("/sandbox/json", func (writer http.ResponseWriter, r *http.Request){
    	var f interface{}
 		body, _ := ioutil.ReadAll(r.Body)
    	json.Unmarshal(body, &f)
    	msg,err := json.MarshalIndent(&f, "", "    ")
    	if err!=nil{
    		stub("RUHROH")
    	}
    	io.WriteString(writer, string(msg))

	})

	http.HandleFunc("/sandbox/jsonspecific", func (writer http.ResponseWriter, r *http.Request){
		type Message struct {
		    Name string
		    Body string
		    Time int64
		}
 		body, _ := ioutil.ReadAll(r.Body)
	    var m Message
	    json.Unmarshal(body, &m)

	    io.WriteString(writer, fmt.Sprintf( r.Method, "...", m.Name))
	})

    http.ListenAndServe(":7776", nil)
}
