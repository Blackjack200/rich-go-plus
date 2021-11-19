// +build windows

package rpc

import (
	"gopkg.in/natefinch/npipe.v2"
	"net"
)

func open() (net.Conn, error) {
	conn, err := npipe.DialTimeout(`\\.\pipe\discord-ipc-0`, timeout)
	if err != nil {
		return conn, err
	}
	return conn, nil
}
