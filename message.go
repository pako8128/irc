package irc

import (
	"strings"
)

type Message struct {
	Prefix  string
	Command string
	Params  []string
}

func (m *Message) Server() string {
	if m.Host() != "" {
		return ""
	}
	return m.Prefix
}

func (m *Message) Nick() string {
	bang := strings.IndexByte(m.Prefix, '!')
	if bang == -1 {
		return ""
	}
	return m.Prefix[:bang]
}

func (m *Message) User() string {
	at := strings.IndexByte(m.Prefix, '@')
	if at == -1 {
		return ""
	}
	bang := strings.IndexByte(m.Prefix[:at], '!')
	if bang == -1 {
		return m.Prefix[:at]
	}
	return m.Prefix[bang+1 : at]
}

func (m *Message) Host() string {
	at := strings.IndexByte(m.Prefix, '@')
	if at == -1 {
		return ""
	}
	return m.Prefix[at+1:]
}

func (m Message) String() string {
	b := strings.Builder{}
	if m.Prefix != "" {
		b.WriteString(":" + m.Prefix + " ")
	}
	b.WriteString(m.Command)
	for i, p := range m.Params {
		b.WriteByte(' ')
		if i == len(m.Params)-1 && strings.Contains(p, " ") {
			b.WriteByte(':')
		}
		b.WriteString(p)
	}
	b.WriteString("\r\n")
	return b.String()
}

func Parse(input string) (*Message, error) {
	prefix, input := parsePrefix(input)
	command, input := parseCommand(input)
	params := parseParams(input)
	return &Message{prefix, command, params}, nil
}

func parsePrefix(input string) (string, string) {
	if !strings.HasPrefix(input, ":") {
		return "", input
	}
	space := strings.IndexByte(input, ' ')
	return input[1:space], input[space+1:]
}

func parseCommand(input string) (string, string) {
	space := strings.IndexByte(input, ' ')
	if space == -1 {
		return input, ""
	}
	return input[:space], input[space+1:]
}

func parseParams(input string) []string {
	parts := strings.SplitN(input, ":", 2)
	params := strings.Split(strings.TrimSpace(parts[0]), " ")
	if params[0] == "" {
		params = []string{}
	}
	if len(parts) > 1 {
		params = append(params, parts[1])
	}
	return params
}
