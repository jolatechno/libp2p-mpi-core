package main

import (
	shell "github.com/ipfs/go-ipfs-api"
)

type MessageShell struct {
	Shell *shell.Shell
}

func (m *MessageShell)Write(data []byte) ([]byte, error) {
	obj := &shell.IpfsObject { Data: string(data) }
	hash, err := m.Shell.ObjectPut(obj)
	return []byte(hash), err
}

func (m *MessageShell)Read(hash []byte) ([]byte, error) {
	obj, err := m.Shell.ObjectGet(string(hash))
	if err != nil {
		return []byte{}, err
	}

	return []byte(obj.Data), nil
}