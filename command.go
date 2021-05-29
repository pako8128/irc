package irc

import "errors"

type command int

const (
	// 3.1 Connection Registration
	PASS command = iota
	NICK
	USER
	OPER
	MODE
	SERVICE
	QUIT
	SQUIT
	// 3.2 Channel Operations
	JOIN
	PART
	TOPIC
	NAMES
	LIST
	INVITE
	KICK
	// 3.3 Sending Messages
	PRIVMSG
	NOTICE
	// 3.4 Server Queries and Commands
	MOTD
	LUSERS
	VERSION
	STATS
	LINKS
	TIME
	CONNECT
	TRACE
	ADMIN
	INFO
	// 3.5 Service Query and Commands
	SERVLIST
	SQUERY
	// 3.6 User based Queries
	WHO
	WHOIS
	WHOWAS
	// 3.7 Miscellaneous Messages
	KILL
	PING
	PONG
	ERROR
	// 4. Optional Features
	AWAY
	REHASH
	DIE
	RESTART
	SUMMON
	USERS
	WALLOPS
	USERHOST
	ISON
)

var cmdMap = map[command]string{
	PASS:     "PASS",
	NICK:     "NICK",
	USER:     "USER",
	OPER:     "OPER",
	MODE:     "MODE",
	SERVICE:  "SERVICE",
	QUIT:     "QUIT",
	SQUIT:    "SQUIT",
	PART:     "PART",
	TOPIC:    "TOPIC",
	NAMES:    "NAMES",
	LIST:     "LIST",
	INVITE:   "INVITE",
	KICK:     "KICK",
	PRIVMSG:  "PRIVMSG",
	NOTICE:   "NOTICE",
	MOTD:     "MOTD",
	LUSERS:   "LUSERS",
	VERSION:  "VERSION",
	STATS:    "STATS",
	LINKS:    "LINKS",
	CONNECT:  "CONNECT",
	TRACE:    "TRACE",
	ADMIN:    "ADMIN",
	INFO:     "INFO",
	SERVLIST: "SERVLIST",
	SQUERY:   "SQUERY",
	WHO:      "WHO",
	WHOIS:    "WHOIS",
	WHOWAS:   "WHOWAS",
	KILL:     "KILL",
	PING:     "PING",
	PONG:     "PONG",
	ERROR:    "ERROR",
	AWAY:     "AWAY",
	REHASH:   "REHASH",
	DIE:      "DIE",
	RESTART:  "RESTART",
	SUMMON:   "SUMMON",
	USERS:    "USERS",
	WALLOPS:  "WALLOPS",
	USERHOST: "USERHOST",
	ISON:     "ISON",
}

func (c command) String() string {
	return cmdMap[c]
}

func parseCmd(cmd string) (command, error) {
	for k, v := range cmdMap {
		if v == cmd {
			return k, nil
		}
	}
	return 0, errors.New("not a valid command")
}
