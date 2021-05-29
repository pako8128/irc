package irc

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Client struct {
	sc *bufio.Scanner
	w  io.Writer
}

func NewClient(rw io.ReadWriter) *Client {
	return &Client{
		sc: bufio.NewScanner(rw),
		w:  rw,
	}
}

func (c *Client) Authenticate(nick, user, pass string) error {
	if pass != "" {
		if err := c.Send(&Message{
			Command: PASS,
			Params:  []string{pass},
		}); err != nil {
			return err
		}
	}

	if err := c.Send(&Message{
		Command: NICK,
		Params:  []string{nick},
	}); err != nil {
		return err
	}

	return c.Send(&Message{
		Command: USER,
		Params:  []string{nick, "8", "*", user},
	})
}

func (c *Client) Join(channels ...string) error {
	return c.Send(&Message{
		Command: JOIN,
		Params:  []string{strings.Join(channels, ",")},
	})
}

func (c *Client) Send(m *Message) error {
	_, err := fmt.Fprint(c.w, m.String())
	return err
}

func (c *Client) Recv() (*Message, error) {
	if !c.sc.Scan() {
		err := c.sc.Err()
		if err == nil {
			return nil, io.EOF
		}
		return nil, err
	}

	return Parse(c.sc.Text())
}
