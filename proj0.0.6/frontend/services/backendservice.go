/*
	Author: Kyle Ong
	Date: 10/25/2018

	backendservice for frontend server
*/
package backendservice

import (
	// standard
	"encoding/gob"
	"fmt"
	"log"
	"net" // user-defined

	"distsys/proj0.0.6/datamodels"
	"distsys/proj0.0.6/vr/flags"
	"distsys/proj0.0.6/vr/globals"
	vrrpc "distsys/proj0.0.6/vr/rpc"
)

type Service interface {
	/*
		service for frontend to interface with backend
	*/
	Get(owner string) []datamodels.Item
	Save(owner string, newItems []datamodels.Item) error
}

type Message struct {
	/*
		messages passed between front-end and back-end
	*/
	SessionID, HttpMethod string
	Body                  []datamodels.Item
}

type DataService struct {
	HostName   string
	RequestNum int
}

func NewDataService(hostName string) *DataService {
	/*
		new DataService for the backend
	*/
	return &DataService{
		HostName:   hostName,
		RequestNum: 0}
}

func (s *DataService) Get(sessionOwner string) []datamodels.Item {
	/*
		GET readingList entries by session
	*/
	message := vrrpc.Message{
		SessionID:  sessionOwner,
		HttpMethod: "GET",
	}

	fmt.Println("this message is about to be sent:", message)

	request := vrrpc.Request{
		Op: vrrpc.Operation{
			Message: message,
		},
		ClientId:   *flags.Listen,
		RequestNum: s.RequestNum,
	}

	fmt.Println("this request is about to be sent:", request)

	conn, err := net.Dial("tcp", s.HostName)
	if err != nil {
		log.Fatal("Connection error", err)
	} else {
		fmt.Println("Connection Established!")
	}

	defer conn.Close()

	encoder := gob.NewEncoder(conn)
	encoder.Encode(request)

	decoder := gob.NewDecoder(conn)
	//payload := &[]datamodels.Item{}
	//will need to update the response type to vrrpc.Response
	resp := &vrrpc.Response{}
	decoder.Decode(resp)

	s.processResp(resp)
	s.RequestNum++

	return resp.OpResult.Message.Body
}

func (s *DataService) Save(sessionOwner string, newItems []datamodels.Item) error {
	/*
		POST a readingList entry
	*/
	conn, err := net.Dial("tcp", s.HostName)
	if err != nil {
		log.Fatal("Connection error", err)
	} else {
		fmt.Println("Connection Established")
	}

	defer conn.Close()

	encoder := gob.NewEncoder(conn)
	message := vrrpc.Message{
		SessionID:  sessionOwner,
		HttpMethod: "SAVE",
		Body:       newItems}

	encoder.Encode(message)

	return nil
}

func (s *DataService) currentPrimaryId() (int, error) {
	for id, p := range globals.AllPorts {
		if p == globals.Port {
			return id, nil
		}
	}
	return 0, fmt.Errorf("cannot find id corresponding to port %v", globals.Port)
}

func (s *DataService) processResp(resp *vrrpc.Response) {
	curID, err := s.currentPrimaryId()
	if err != nil {
		log.Fatalf("Failed to look up primary ID: %v", err)
	}
	log.Printf("current view num: %v", resp.ViewNum)

	if errMsg := resp.Err; errMsg != "" {
		if errMsg == "not primary" {
			newId := resp.ViewNum % len(globals.AllPorts)
			log.Printf("Primary %v => %v", curID, newId)
			globals.Port = globals.AllPorts[newId]
		} else if errMsg == "view change" {
			log.Printf("currently under view change")
		} else {
			log.Printf("got error message but it was not recognized: %v", errMsg)
		}
		return
	}

	fmt.Printf("Vrgo response: %v\n", resp.OpResult.Message)
}
