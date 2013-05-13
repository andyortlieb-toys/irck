package main

import (
	//"github.com/bmizerany/pq"
	"fmt"
)

func stub(msg string) {
	fmt.Println("STUB: ", msg)
}
func erro(msg string) {
	fmt.Println("ERROR: ", msg)
}

func initUsers() []*User {

	stub("Set up some users")

	guyA := User{username: "Acdtrux"}

	
	idtA1 := guyA.AddIdentity(
		"Freenode",
		"acidTrucks",
		"irc",
		"irc.freenode.net:6667",
		true,
	)
	//idtA1.AddChannel("#botters", true)
	idtA1.AddChannel("#devmke", true)
	idtA1.Connect()
	
	
	idtA2 := guyA.AddIdentity(
		"corvisa",
		"AndyO",
		"irc",
		"dv.opasc.net:6667",
		true,
	)

	idtA2.AddChannel("#dev", true)
	idtA2.AddChannel("#advent-dev", true)
	idtA2.AddChannel("#cloud", true)
	idtA2.AddChannel("#qc", true)
	idtA2.Connect()
	
	
	idtA3 := guyA.AddIdentity(
		"dv-2",
		"Acdtrux",
		"irc",
		"dv.opasc.net:6667",
		true,
	)

	idtA3.AddChannel("#dingolove", true)
	idtA3.AddChannel("#blabla", true)
	idtA3.Connect()

	guyB := User{username: "dingolvr"}
	idtB1 := guyB.AddIdentity(
		"dv",
		"dingolvr",
		"irc",
		"dv.opasc.net:6667",
		true,
	)
	idtB1.AddChannel("#dingolove", true)
	idtB1.AddChannel("#yadda", true)
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
	return []*User{
		&guyA,
		&guyB,
	}
}

func main() {
	stub("Starting irckd")
	initHttp(initUsers())
}
