/*
	Author: Kyle Ong
	Date: 10/25/2018

	test suite for utils
*/
package utils

import (
	"fmt"
	"testing"
)

func TestParseListenPort(t *testing.T) {
	/*
		test parse listen port
	*/
	args := make([]string, 4)
	args[0] = "--listen"
	args[1] = "8899"
	args[2] = "--backend"
	args[3] = "hostName:8090"
	portNum := ParseListenPort(args)
	if portNum != "8899" {
		message := fmt.Sprintf("Expected: 8899. Got: %s", portNum)
		t.Error(message)
	}
}

func TestParseBackendPort(t *testing.T) {
	/*
		tests parse backend port
	*/
	args := make([]string, 4)
	args[0] = "--listen"
	args[1] = "8899"
	args[2] = "--backend"
	args[3] = "hostName:8090"
	portNum := ParseBackendHost(args)
	if portNum != "hostName:8090" {
		message := fmt.Sprintf("Expected: 'hostName:8090'. Got: %s", portNum)
		t.Error(message)
	}
}
