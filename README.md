
# Irck! The IRC Keeper

Irck is a persistent irc client-daemon service.  It can handle man users, each on many networks.  It currently provides a web-based client, and one of the highest priority goals is to provide an IRC server-based client, so you can connect to Irck with whichever IRC client you are most comfortable--and you can optionally make only one client connection for your many service connections!  But that's not here yet!

Check out TODO.md if you're interested in development.

## System requirements

### irckd:
An architecture for wish a release build is available, or else a compiler for the go programming language.

### Web client:
Chrome only.  Maybe others, but you're on your own.  

## Upcoming Goals:

    * Mobile client (android.  Any takers for iphone?)
    * IRCK console (prvmsg, with repl-like ui)
    * IRC Server-based client
        * Connect to one specific identity
        * Connect to all identities through one IRC client
        * Autoscroll only when sane.
    * Web-based client
        * Name highlighting
            * Bubble up through buttons (channel/identity)
        * Inline image/video/audio dereferencing (by preference)
    * DCC support:
        * Back end to automatically accept & hold files
            * For configurable period
            * Until told by user to remove
            * By user's preference
        * Web-based client 
            * Just a list of files, where they came from, etc.  click to download.
        * IRC server client:
            * Through the irck console
    * Automatically handle messages
        * Condition support:
            * When inactive for X-time
            * When X-clients are connected/active
        * Auto respond (away)
        * SMS notification
        * Email notification

## Moving parts

* Data
	* User
		* Auth stuff/ profile
		* Identity
			* Service type
			* Service info/profile

* Libraries, Dependencies
	* "github.com/thoj/go-ircevent"
	* 
