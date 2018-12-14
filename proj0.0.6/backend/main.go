/*
	Author: Kyle Ong
	Date: 10/25/2018

	backend server for readinglist application
*/
package main

import (
	//standard libary
	//user defined
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"

	"distsys/proj0.0.6/backend/datasource"
	"distsys/proj0.0.6/backend/services"
	"distsys/proj0.0.6/backend/utils"
	"distsys/proj0.0.6/datamodels"
	"distsys/proj0.0.6/vr/globals"
	"distsys/proj0.0.6/vr/oplog"
	vrrpc "distsys/proj0.0.6/vr/rpc"
	"distsys/proj0.0.6/vr/table"
	cache "github.com/patrickmn/go-cache"
)

type ClientRequest struct {
	Request vrrpc.Request
	done    chan *vrrpc.Response
}

type AckMessage struct{ Status int }

func execute(req *vrrpc.Request, resp *vrrpc.Response) error {
	mode := globals.Mode

	if mode != "primary" {
		globals.Log("Execute", "not primary; view num: %v", globals.ViewNum)
		var err string
		if mode == "backup" {
			globals.Log("Execute", "I am not primary anymore; view num: %v", globals.ViewNum)
			err = fmt.Sprintf("not primary")
		} else if mode == "viewchange" || mode == "viewchange-init" {
			globals.Log("Execute", "under view change")
			err = fmt.Sprintf("view change")
		}
		*resp = vrrpc.Response{
			ViewNum: globals.ViewNum,
			Err:     err,
		}
		return nil
	}
	return nil
}

func handleConnection(conn net.Conn, dataService *service.DataService) {
	/*
		process incoming requests
	*/
	defer conn.Close()

	decoder := gob.NewDecoder(conn)
	clientRequest := &vrrpc.Request{}
	decoder.Decode(clientRequest)
	fmt.Println("Recieved Client Request:", clientRequest)
	message := vrrpc.Message{}
	message = clientRequest.Op.Message
	fmt.Println("Received Client Message:", message)
	encoder := gob.NewEncoder(conn)
	body := []datamodels.Item{}
	resp := vrrpc.Response{}
	if message.HttpMethod == "GET" {
		body = dataService.Get(message.SessionID)
		resp = vrrpc.Response{
			OpResult: vrrpc.OperationResult{
				Message: vrrpc.Message{Body: body}}}
		execute(clientRequest, &resp)
		encoder.Encode(resp)
	} else if message.HttpMethod == "SAVE" {
		err := dataService.Save(message.SessionID, message.Body)
		if err != nil {
			log.Fatal(err)
		}
		body = dataService.Get(message.SessionID)
		resp = vrrpc.Response{
			OpResult: vrrpc.OperationResult{
				Message: vrrpc.Message{Body: body}}}
		execute(clientRequest, &resp)
		encoder.Encode(body)
	} else if message.HttpMethod == "PING" {
		payload := AckMessage{
			Status: 1}
		encoder.Encode(payload)
	}

}

func main() {
	/*
		entry point for reading list
	*/
	//ctx := context.Background()
	globals.ClientTable = table.New(cache.NoExpiration, cache.NoExpiration)
	globals.OpLog = oplog.New()

	dataService := service.NewDataService(datasource.Items)
	args := os.Args[1:]
	portNum := utils.ParseListenPort(args)
	replicas := utils.ParseReplicas(args)
	fmt.Println("Replicas:", replicas)
	fmt.Println("All Ports:", globals.AllPorts)
	fmt.Println("All Other Ports:", globals.AllOtherPorts())
	fmt.Println("Starting backend server...")
	ln, err := net.Listen("tcp", ":"+portNum)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("listening on port:", portNum)
	for {
		conn, err := ln.Accept() // this blocks until connection or error
		if err != nil {
			log.Fatal("Setup error", err)
		}
		handleConnection(conn, dataService) // a goroutine handles conn so that the loop can accept other connections
	}
}
