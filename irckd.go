package main

import (
	"github.com/thoj/go-ircevent"
	"fmt"
)


func Connect() *irc.Connection {
	//irccon := IRC("go-eventirc", "go-eventirc")
	irccon := irc.IRC("dingolvr", "dingolvr")
	irccon.VerboseCallbackHandler = true
	err := irccon.Connect("irc.freenode.net:6667")
	if err != nil {
		fmt.Println("Can't connect to freenode.")
	}
	return irccon
}

func Manage(irccon *irc.Connection) {
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join("#dingolove") })

	irccon.AddCallback("366", func(e *irc.Event) {
		irccon.Privmsg("#dingolove", "Test Message\n")
		irccon.Nick("dingolvr_newnick")
	})

	irccon.Loop()
}

func main(){
	irccon := Connect();
	var input string

	go Manage(irccon)

	for {
		fmt.Scanln(&input)
		if input == "/quit" {
			irccon.Quit()
			return
		}
		irccon.Privmsg("#dingolove", input)
	}

}