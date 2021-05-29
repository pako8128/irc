package irc

import (
	"testing"
)

func compareMessage(m1, m2 *Message) bool {
	if m1.Command != m2.Command ||
		m1.Prefix != m2.Prefix ||
		len(m1.Params) != len(m2.Params) {
		return false
	}

	for i := range m1.Params {
		if m1.Params[i] != m2.Params[i] {
			return false
		}
	}

	return true
}

func TestParseMessage(t *testing.T) {
	tt := []struct {
		raw string
		mes Message
		err error
	}{
		{
			":syrk!kalt@millennium.stealth.net QUIT :Gone to have lunch",
			Message{
				Prefix:  "syrk!kalt@millennium.stealth.net",
				Command: QUIT,
				Params:  []string{"Gone to have lunch"},
			},
			nil,
		},
		{
			":Trillian SQUIT cm22.eng.umd.edu :Server out of control",
			Message{
				Prefix:  "Trillian",
				Command: SQUIT,
				Params:  []string{"cm22.eng.umd.edu", "Server out of control"},
			},
			nil,
		},
		{
			":WiZ!jto@tolsun.oulu.fi JOIN #Twilight_zone",
			Message{
				Prefix:  "WiZ!jto@tolsun.oulu.fi",
				Command: JOIN,
				Params:  []string{"#Twilight_zone"},
			},
			nil,
		},
		{
			":WiZ!jto@tolsun.oulu.fi PART #playzone :I lost",
			Message{
				Prefix:  "WiZ!jto@tolsun.oulu.fi",
				Command: PART,
				Params:  []string{"#playzone", "I lost"},
			},
			nil,
		},
		{
			":WiZ!jto@tolsun.oulu.fi MODE #eu-opers -l",
			Message{
				Prefix:  "WiZ!jto@tolsun.oulu.fi",
				Command: MODE,
				Params:  []string{"#eu-opers", "-l"},
			},
			nil,
		},
		{
			"MOTD",
			Message{
				Command: MOTD,
			},
			nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.mes.Command, func(t *testing.T) {
			t.Parallel()
			parsed, err := Parse(tc.raw)
			if err != tc.err {
				t.Fatalf("error: expected: %v\ngot: %v", tc.err, err)
			}
			if !compareMessage(parsed, &tc.mes) {
				t.Errorf("expected: %+v\ngot: %+v", tc.mes, parsed)
			}
			if tc.mes.String() != tc.raw+"\r\n" {
				t.Errorf("expected: %+v\ngot: %+v", tc.raw, tc.mes.String())
			}
		})
	}
}

func TestPrefix(t *testing.T) {
	tt := []struct {
		message *Message
		server  string
		nick    string
		user    string
		host    string
	}{
		{
			message: &Message{Prefix: "WiZ!jto@tolsun.oulu.fi"},
			nick:    "WiZ",
			user:    "jto",
			host:    "tolsun.oulu.fi",
		},
		{
			message: &Message{Prefix: "syrk!kalt@millennium.stealth.net"},
			nick:    "syrk",
			user:    "kalt",
			host:    "millennium.stealth.net",
		},
		{
			message: &Message{Prefix: "Trillian"},
			server:  "Trillian",
		},
		{
			message: &Message{Prefix: ""},
		},
	}

	for _, tc := range tt {
		t.Run(tc.message.Prefix, func(t *testing.T) {
			t.Parallel()
			if tc.message.Server() != tc.server {
				t.Errorf("wrong server: expected %q got %q", tc.server, tc.message.Server())
			}
			if tc.message.Nick() != tc.nick {
				t.Errorf("wrong nick: expected %q got %q", tc.nick, tc.message.Nick())
			}
			if tc.message.User() != tc.user {
				t.Errorf("wrong user: expected %q got %q", tc.user, tc.message.User())
			}
			if tc.message.Host() != tc.host {
				t.Errorf("wrong server: expected %q got %q", tc.host, tc.message.Host())
			}
		})
	}
}
