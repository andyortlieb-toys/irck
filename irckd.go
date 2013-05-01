package main

import (
	"github.com/thoj/go-ircevent"
	//"github.com/bmizerany/pq"
	"fmt"
	"bufio"
	"os"
)

type Channel struct {
	name		string
	history 	[]string
}

type Identity struct {
	servertype	string
	serveraddr	string
	username	string
	channels	[]Channel
	connection	irc.Connection
}

type User struct {
	username	string
	identities	[]Identity
}

func Create() *irc.Connection {
	//irccon := IRC("go-eventirc", "go-eventirc")
	irccon := irc.IRC("dingolvr", "dingolvr")
	irccon.VerboseCallbackHandler = true
	return irccon
}

func Connect(irccon *irc.Connection) {
	// Connect
	err := irccon.Connect("irc.freenode.net:6667")
	if err != nil {
		fmt.Println("Can't connect to freenode.")
	}

	// Join all of the channels.
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join("#dingolove") })
	
	// Hang out.
	irccon.Loop()
}

func main(){
	br := bufio.NewReaderSize(os.Stdin, 512)
	irccon := Create();

	go Connect(irccon)

	for {
		msg, err := br.ReadString('\n')
		if err != nil {
			irccon.Quit()
			break;
		}
		if msg[:5] == "/quit" {
			irccon.Quit()
			return
		}
		fmt.Println("#dingolove/dingolvr: ", msg)
		irccon.Privmsg("#dingolove", msg+"\n")
	}

}