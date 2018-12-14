package utils

import(
	"strings"
	"bytes"
)

func ParseListenPort(args []string) string {
	portNum := "8080"
	for idx, ele := range args {
		if ele == "--listen" && idx < (len(args)-1) {
			portNum = args[idx+1]
		}
	}
	return portNum
}

func ParseBackendPort(args []string) string{
	backend := "localhost:8090"
	for idx, ele := range args{
		if ele == "--backend" &&  idx < len(args) - 1{
			backend = args[idx + 1]
		}
	}
	if len(strings.Split(backend, ":")) == 1{
		var buffer bytes.Buffer
		buffer.WriteString("localhost")
		buffer.WriteString(backend)
		backend = buffer.String()
	}
	return backend
}
