package main

import (
	"github.com/thoj/go-ircevent"
	"net/http"
	//"github.com/bmizerany/pq"
	"fmt"
	//"time"
	"io"
	"io/ioutil"
	// "bufio"
	// "os"
	"encoding/json"
)

func stub(msg string){
	fmt.Println("STUB: ", msg)
}
func erro(msg string){
	fmt.Println("ERROR: ", msg)
}

type User struct {
	username	string
	identities	[]*Identity
}

type Identity struct {
	servertype	string
	serveraddr	string
	nick	string
	enabled		bool
	channels	[]Channel
	connection	*irc.Connection
	user 		*User

}

type Channel struct {
	name		string
	history 	*[]string
}

func (usr *User) AddIdentity(stype string,addr string,nick string,enabled bool) *Identity{
	idt := Identity{
		servertype:stype,
		serveraddr:addr,
		nick:nick,
		enabled:enabled,
		user:usr,
	}
	usr.identities = append(usr.identities, &idt)
	return &idt
}

func (idt *Identity) JoinChannels(){
	stub("Joining Channels")
	idt.connection.Join("#dingolove")
}

func (idt *Identity) Connect() *irc.Connection{
	// Initialize a connection
	irccon := irc.IRC(idt.nick, idt.user.username)
	irccon.VerboseCallbackHandler = true
	idt.connection = irccon

	// Really connect
	err := irccon.Connect(idt.serveraddr)
	if err != nil{
		erro(fmt.Sprintf("Cannot connect to `", idt.serveraddr, "` for user `", idt.user.username))
	}

	// Manage it
	irccon.AddCallback("001", func(e *irc.Event){idt.JoinChannels()})
	go func(){
		irccon.Loop()
	}()
	return irccon
}

func (idt *Identity) AddChannel(msg string){
	cn := Channel{
		name:msg,
	}
	idt.channels = append(idt.channels, cn)
	stub("TODO: Connect to channel if we are already connected to the host")
}

func initUsers() *[]User{

	stub("Set up some users")
	guyA := User{username: "dingolvrA"}
	idtA1 := guyA.AddIdentity(
		"irc",
		"dv.opasc.net:6667",
		"dingolvrA1",
		true,
	)
	idtA2 := guyA.AddIdentity(
		"irc",
		"dv.opasc.net:6667",
		"dingolvrA2",
		true,
	)
	fmt.Println(20, idtA1.channels, guyA.identities[0].channels)
	idtA1.AddChannel("#dingolove")
	fmt.Println(21, idtA1.channels, guyA.identities[0].channels)
	idtA2.AddChannel("#dingolove")
	idtA1.Connect()
	idtA2.Connect()

	guyB := User{username: "dingolvrB"}
	idtB1 := guyB.AddIdentity(
		"irc",
		"dv.opasc.net:6667",
		"dingolvrB1",
		true,
	)
	idtB1.AddChannel("#dingolove")
	idtB1.Connect()

/*	// Just for now... loop with readstring until we know how to be a real good daemon.

	br := bufio.NewReaderSize(os.Stdin, 512)
	for {
		msg, err := br.ReadString('\n')
		if err != nil {
			break;
		}
		if msg[:5] == "/quit" {
			return
		}
		fmt.Println("#dingolove/dingolvr: ", msg)
		idtA1.connection.Privmsg("AcidTrucks", msg+"\n")
		idtA2.connection.Privmsg("#dingolove", msg+"\n")
		//irccon.Privmsg("#dingolove", msg+"\n")
	}	
*/
	return &[]User{
		guyA,
		guyB,
	}
}

func initHttp(users *[]User){

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


    http.HandleFunc("/msg/", func (writer http.ResponseWriter, r *http.Request){
		io.WriteString(writer, fmt.Sprintf("msg!",writer,"\n"))
	})

    http.HandleFunc("/history/", func (writer http.ResponseWriter, r *http.Request){
		io.WriteString(writer, "history!\n")
	})

    http.ListenAndServe(":7776", nil)
}

func main(){
	stub("Starting irckd")
	initHttp(initUsers())
}

