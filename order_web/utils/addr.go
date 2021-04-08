package utils

import (
	"fmt"
	"net"
)

func GetFreePort() (int, error) {
	address, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:0", "0.0.0.0"))
	if err != nil {
		return 0, err
	}
	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		return 0, err
	}
	port := listener.Addr().(*net.TCPAddr).Port
	//先关闭return
	_ = listener.Close()
	return port, nil

}
