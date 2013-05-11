package main

import (
	"fmt"
	"time"
	"sync"
	"github.com/thoj/go-ircevent"
)

type User struct {
	Identities []*Identity
	HistoryIdx	int

	username   string
	password   string
	identityIdx	int
}

func (usr *User) IdentityIncr() int {
	usr.identityIdx++
	return usr.identityIdx
}

func (usr *User) HistoryIncr() int {
	usr.HistoryIdx++
	return usr.HistoryIdx
}

func (usr *User) AddIdentity(sname string, nick string, stype string, addr string, enabled bool) *Identity {
	idt := Identity{
		IdentityIdx: usr.IdentityIncr(),
		Servername: sname,
		Nick:       nick,

		servertype: stype,
		serveraddr: addr,
		enabled:    enabled,
		user:       usr,
	}
	usr.Identities = append(usr.Identities, &idt)
	return &idt
}

type History struct {
	Time		time.Time
	Originator 	string
	Recipient 	string
	Message		string
	Raw 		string
	HistoryIdx	int
	IdentityIdx int

	event		*irc.Event
}

type Identity struct {
	IdentityIdx	int
	Servername string
	Nick       string
	History 	[]*History

	servertype string
	serveraddr string
	enabled    bool
	channels   []*Channel
	connection *irc.Connection
	user       *User
	watchers    map[*func(*History)] *func(*History)
	watcherctl  sync.WaitGroup
}

func (idt *Identity) AddWatcher(fn *func(*History)) {
	// Initialize my watcher controls if needed.
	if idt.watchers == nil { idt.watchers = make( map[*func(*History)] *func(*History) )}

	idt.watchers[fn] = fn
}

func (idt *Identity) RemoveWatcher(fn *func(*History)) {
	delete (idt.watchers, fn)
}

func (idt *Identity) JoinChannels() {
	stub("Joining Channels")

	/*

				FIXME !!!!!

	*/

	idt.connection.Join("#dingolove")
}

func (idt *Identity) Connect() *irc.Connection {
	// Initialize a connection
	irccon := irc.IRC(idt.Nick, idt.user.username)
	irccon.VerboseCallbackHandler = true
	idt.connection = irccon

	// Really connect
	err := irccon.Connect(idt.serveraddr)
	if err != nil {
		erro(fmt.Sprintf("Cannot connect to `", idt.serveraddr, "` for user `", idt.user.username))
	}

	// Manage it
	irccon.AddCallback("001", func(e *irc.Event) { idt.JoinChannels() })

	historyCallback := func(e *irc.Event){
		// Create the history instance
		hst := History{}
		hst.event = e
		hst.Originator = e.Nick
		hst.Message = e.Message
		hst.Recipient = e.Arguments[0]
		hst.Time = time.Now()
		hst.Raw = e.Raw
		hst.HistoryIdx = idt.user.HistoryIncr()
		hst.IdentityIdx = idt.IdentityIdx

		idt.AddHistory(&hst)
	}

	irccon.AddCallback("PRIVMSG", historyCallback)
	/*
	irccon.AddCallback("NOTICE", historyCallback)
	irccon.AddCallback("JOIN", historyCallback)
	*/
	go func() {
		irccon.Loop()
	}()
	return irccon
}

func (idt *Identity) AddHistory(hst *History){
	hst.HistoryIdx = idt.user.HistoryIncr()
	hst.IdentityIdx = idt.IdentityIdx
	idt.History = append(idt.History, hst)
	idt.RunWatchers(hst)
}

func (idt *Identity) RunWatchers(hst *History){
	// Run the watchers
	for _,w := range idt.watchers{
		(*w)(hst)
	}
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
