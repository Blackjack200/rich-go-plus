package rpc

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/blackjack200/rich-go-plus/codec"
	"net"
	"sync"
	"time"
	"unsafe"
)

const timeout = time.Second * 2

type Conn struct {
	c    net.Conn
	once sync.Once
}

func (c *Conn) Close() {
	c.once.Do(func() {
		_ = c.Send(codec.OpcodeClose, "{}")
		_ = c.c.Close()
	})
}

func (c *Conn) Read() (*codec.Message, bool) {
	buf := make([]byte, 512)
	l, err := c.c.Read(buf)
	if err != nil {
		return nil, false
	}
	b := buf[:l]
	hdr := (*codec.MessageFrameHeader)(unsafe.Pointer(&b[0]))
	return &codec.Message{
		Hdr:     hdr,
		Payload: string(b[unsafe.Sizeof(*hdr):]),
	}, true
}

func (c *Conn) Send(opcode int, payload string) error {
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.LittleEndian, int32(opcode)); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.LittleEndian, int32(len(payload))); err != nil {
		return err
	}
	buf.Write([]byte(payload))
	if _, err := c.c.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func Login(id string) (*Conn, error) {
	if payload, err := json.Marshal(codec.HandshakeRequest{
		V:        "1",
		ClientId: id,
	}); err != nil {
		return nil, err
	} else if conn, err := open(); err != nil {
		return nil, err
	} else {
		c := &Conn{
			c: conn,
		}
		if err := c.Send(codec.OpcodeHandshake, string(payload)); err != nil {
			return nil, err
		}
		if m, success := c.Read(); success != true {
			return c, fmt.Errorf("failed to read packet")
		} else {
			if !m.Success() {
				return nil, fmt.Errorf("failed to handshake opcode: %v msg:%v", m.Payload[0], m.Payload)
			}
			return c, nil
		}
	}
}
