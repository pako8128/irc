package irc

type Message struct {
	prefix  string
	command command
	args    []string
}
