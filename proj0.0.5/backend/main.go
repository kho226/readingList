/*
	Author: Kyle Ong
	Date: 10/25/2018

	backend server for readinglist application
*/
package main

import (
	//standard libary
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os" //user defined

	"distsys/proj0.0.5/backend/datamodels"
	"distsys/proj0.0.5/backend/datasource"
	"distsys/proj0.0.5/backend/services"
	"distsys/proj0.0.5/backend/utils"
)

type Message struct {
	SessionID, HttpMethod string
	Body                  []datamodels.Item
}

type AckMessage struct{ Status int }

func handleConnection(conn net.Conn, dataService *service.DataService) {
	defer conn.Close()
	decoder := gob.NewDecoder(conn)
	message := &Message{}
	decoder.Decode(message)
	fmt.Printf("Received : %+v", message)
	encoder := gob.NewEncoder(conn)
	payload := []datamodels.Item{}
	if message.HttpMethod == "GET" {
		payload = dataService.Get(message.SessionID)
		encoder.Encode(payload)
	} else if message.HttpMethod == "SAVE" {
		err := dataService.Save(message.SessionID, message.Body)
		if err != nil {
			log.Fatal(err)
		}
		payload = dataService.Get(message.SessionID)
		encoder.Encode(payload)
	} else if message.HttpMethod == "PING" {
		payload := AckMessage{
			Status: 1}
		encoder.Encode(payload)
	}

}

func main() {
	dataService := service.NewDataService(datasource.Items)
	args := os.Args[1:]
	portNum := utils.ParseListenPort(args)
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
