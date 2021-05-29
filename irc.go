package irc

import (
	"bufio"
	"fmt"
	"io"
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
