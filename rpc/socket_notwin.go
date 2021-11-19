//+build !windows

package rpc

import (
	"net"
	"os"
	"path/filepath"
)

func open() (net.Conn, error) {
	conn, err := net.DialTimeout("unix", filepath.Join(os.TempDir(), "discord-ipc-0"), timeout)
	if err != nil {
		return conn, err
	}
	return conn, nil
}
