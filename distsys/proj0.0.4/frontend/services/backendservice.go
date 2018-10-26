
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

	"distsys/proj0.0.4/frontend/datamodels"
)

type Service interface {
	Get(owner string) []datamodels.Item
	Save(owner string, newItems []datamodels.Item) error
}

type Message struct {
	SessionID, HttpMethod string
	Body                  []datamodels.Item
}

type DataService struct {
	HostName string
}

func NewDataService(hostName string) *DataService {
	return &DataService{
		HostName: hostName}
}

func (s *DataService) Get(sessionOwner string) []datamodels.Item {
	message := Message{
		SessionID:  sessionOwner,
		HttpMethod: "GET",
	}

	conn, err := net.Dial("tcp", s.HostName)
	if err != nil {
		log.Fatal("Connection error", err)
	} else {
		fmt.Println("Connection Established!")
	}

	defer conn.Close()

	encoder := gob.NewEncoder(conn)
	encoder.Encode(message)

	decoder := gob.NewDecoder(conn)
	payload := &[]datamodels.Item{}
	decoder.Decode(payload)

	return *payload
}

func (s *DataService) Save(sessionOwner string, newItems []datamodels.Item) error {
	conn, err := net.Dial("tcp", s.HostName)
	if err != nil {
		log.Fatal("Connection error", err)
	} else {
		fmt.Println("Connection Established")
	}

	defer conn.Close()

	encoder := gob.NewEncoder(conn)
	message := Message{
		SessionID:  sessionOwner,
		HttpMethod: "SAVE",
		Body:       newItems}
	encoder.Encode(message)

	return nil
}
