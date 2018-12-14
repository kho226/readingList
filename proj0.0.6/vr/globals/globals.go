// globals defines the global variables shared between primary and backup.
package globals

import (
	"context"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"distsys/proj0.0.6/backend/utils"
	"distsys/proj0.0.6/vr/flags"
	"distsys/proj0.0.6/vr/oplog"
	"distsys/proj0.0.6/vr/table"
)

type MutexInt struct {
	/*
		MutexInt is a thread-safe int.
	*/
	sync.Mutex
	V int
}

func (m *MutexInt) Locked(f func()) {
	/*
		Locked locks the int.
	*/
	m.Lock()
	defer m.Unlock()
	f()
}

type MutexString struct {
	/*
		MutexString is a thread-safe string.
	*/
	sync.Mutex
	V string
}

func (m *MutexString) Locked(f func()) {
	/*
	 Locked locks the string.
	*/
	m.Lock()
	defer m.Unlock()
	f()
}

type MutexBool struct {
	/*
		MutexBool is a thread-safe bool.
	*/
	sync.Mutex
	V bool
}

func (m *MutexBool) Locked(f func()) {
	/*
		Locked locks the bool.
	*/
	m.Lock()
	defer m.Unlock()
	f()
}

func Log(f, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Printf("[%v, %20v] %v", *flags.Id, f, msg)
}

var (
	// The port of the replica.
	Port int

	// The Operation request ID.
	OpNum int

	// The current view number.
	ViewNum int

	// The current commit number.
	CommitNum int

	// The mode of the replica. Only monitor is supposed to change this.
	Mode string

	// The operation log.
	OpLog *oplog.OpRequestLog

	// The client table.
	ClientTable *table.ClientTable

	// The global cancellable context.
	CtxCancel context.Context

	// AllPorts is a map from id to port.
	AllPorts = map[int]int{}

	// clients is a map from hostname to *rpc.Client.
	// This way each node only creates one outgoing client to another node,
	// and more requests to the same node will reuse the same client.
	clients = map[string]*rpc.Client{}
)

func init() {
	/*
		init global state for a replica
	*/
	replicas := utils.ParseReplicas(os.Args[1:])
	utils.ParseListenPort(os.Args[1:])
	ports := []int{}
	for _, node := range replicas {
		host := strings.Split(node, ":")
		port, _ := strconv.Atoi(host[1]) //assuming all input is valid
		ports = append(ports, port)
	}
	thisPort, _ := strconv.Atoi(utils.ParseListenPort(os.Args[1:]))
	ports = append(ports, thisPort)
	sort.Ints(ports[:])
	for idx, port := range ports {
		if idx == *flags.Id {
			if idx == 0 && port == thisPort {
				Mode = "primary"
			} else {
				Mode = "backup"
			}
			Port = *flags.Listen
			Log("globals.init", "initial mode: %v; port: %v", Mode, Port)
		}
		AllPorts[idx] = port
	}
}

func AllOtherPorts() []int {
	/*
		AllOtherPorts returns all the other replica ports except for that of the current node.
	*/
	var ps []int
	for _, p := range AllPorts {
		if p != Port {
			ps = append(ps, p)
		}
	}
	return ps
}

func GetOrCreateClient(hostname string) (*rpc.Client, error) {
	/*
		GetOrCreateClient returns a cached rpc.Client or creates a new rpc.Client.
	*/
	if client, ok := clients[hostname]; ok == true {
		return client, nil
	}
	client, err := rpc.DialHTTP("tcp", hostname)
	if err != nil {
		return nil, fmt.Errorf("failed to dial %v: %v", hostname, err)
	}
	clients[hostname] = client
	return client, nil
}
