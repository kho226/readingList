/*
	Author: Kyle Ong
	Date: 10/25/2018

	utilities for backend server
*/

package utils

import (
	"bytes"
	"strings"
)

func ParseListenPort(args []string) string {
	/*
		parse listen port from command line args
	*/
	portNum := "8090"
	for idx, ele := range args {
		if ele == "--listen" && idx < (len(args)-1) {
			portNum = args[idx+1]
		}
	}
	return portNum
}

func ParseBackendHost(args []string) string {
	/*
		parse backend host from command line args
	*/
	backend := "localhost:8090"
	for idx, ele := range args {
		if ele == "--backend" && idx < len(args)-1 {
			backend = args[idx+1]
		}
	}
	if len(strings.Split(backend, ":")) == 1 {
		var buffer bytes.Buffer
		buffer.WriteString("localhost")
		buffer.WriteString(backend)
		backend = buffer.String()
	}
	return backend
}

func ParseReplicas(args []string) []string {
	/*
		parse replica host names from command line args
		assume that input args are valid
		to-do:
			- error checking
	*/
	replicas := []string{}
	for idx, ele := range args {
		if ele == "--backend" {
			replicas = strings.Split(args[idx+1], ",")
		}
	}
	for idx, portNum := range replicas {
		split := strings.Split(portNum, ":")
		joined := strings.Join(split, "")
		if len(joined) == 4 {
			var buffer bytes.Buffer
			buffer.WriteString("localhost")
			buffer.WriteString(portNum)
			replicas[idx] = buffer.String()
		}
	}
	return replicas
}
