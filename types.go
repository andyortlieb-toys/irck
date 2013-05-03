package main

import (
	"fmt"
	"github.com/thoj/go-ircevent"
)

type User struct {
	username   string
	password   string
	identities []*Identity
}

func (usr *User) AddIdentity(stype string, addr string, nick string, enabled bool) *Identity {
	idt := Identity{
		servertype: stype,
		serveraddr: addr,
		nick:       nick,
		enabled:    enabled,
		user:       usr,
	}
	usr.identities = append(usr.identities, &idt)
	return &idt
}

type Identity struct {
	servertype string
	serveraddr string
	nick       string
	enabled    bool
	channels   []*Channel
	connection *irc.Connection
	user       *User
}

func (idt *Identity) JoinChannels() {
	stub("Joining Channels")
	idt.connection.Join("#dingolove")
}

func (idt *Identity) Connect() *irc.Connection {
	// Initialize a connection
	irccon := irc.IRC(idt.nick, idt.user.username)
	irccon.VerboseCallbackHandler = true
	idt.connection = irccon

	// Really connect
	err := irccon.Connect(idt.serveraddr)
	if err != nil {
		erro(fmt.Sprintf("Cannot connect to `", idt.serveraddr, "` for user `", idt.user.username))
	}

	// Manage it
	irccon.AddCallback("001", func(e *irc.Event) { idt.JoinChannels() })
	go func() {
		irccon.Loop()
	}()
	return irccon
}

func (idt *Identity) AddChannel(msg string, enabled bool) {
	cn := Channel{
		name:    msg,
		enabled: enabled,
	}
	idt.channels = append(idt.channels, &cn)
	stub("TODO: Connect to channel if we are already connected to the host")
}

type Channel struct {
	name    string
	history []string
	enabled bool
}
