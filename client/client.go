package client

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/blackjack200/rich-go-plus/codec"
	"github.com/blackjack200/rich-go-plus/rpc"
	"os"
)

type Client struct {
	c *rpc.Conn
}

func Dial(id string) (*Client, error) {
	if c, err := rpc.Login(id); err != nil {
		return nil, err
	} else {
		return &Client{c}, nil
	}
}

func (c *Client) SetActivity(a *codec.Activity) error {
	payload, err := json.Marshal(codec.Frame{
		Cmd: "SET_ACTIVITY",
		Args: codec.Args{
			Pid:      os.Getpid(),
			Activity: codec.MapActivity(a),
		},
		Nonce: getNonce(),
	})
	if err != nil {
		return nil
	}
	return c.c.Send(codec.OpcodeFrame, string(payload))
}

func getNonce() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		return ""
	}
	buf[6] = (buf[6] & 0x0f) | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:])
}

func (c *Client) Read() (*codec.Message, bool) {
	return c.c.Read()
}

func (c *Client) Close() {
	c.c.Close()
}
