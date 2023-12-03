package utils

import (
	"fmt"
	"net"
	"strconv"
)

func CheckPortAvailability(port int) bool {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return false
	}
	defer listener.Close()
	return true
}

func FindAvailablePort(port int) (int, error) {
	for startPort := port; startPort < 65535; startPort++ {
		fmt.Println("checking port: ", startPort)
		if CheckPortAvailability(startPort) {
			return startPort, nil
		}
	}

	return 0, fmt.Errorf("no available port")
}
