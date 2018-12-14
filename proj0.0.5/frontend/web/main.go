/*
	Author: Kyle Ong
	Date: 10/25/2018

	server for front end of readinglist application
*/
package main

import (
	"encoding/gob" //standard
	"fmt"
	"log"
	"net"
	"os"
	"time" // user-defined

	"distsys/proj0.0.5/frontend/datamodels"
	"distsys/proj0.0.5/frontend/services"
	"distsys/proj0.0.5/frontend/utils"
	"distsys/proj0.0.5/frontend/web/controllers"
	"github.com/kataras/iris" // third party
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/websocket"
)

type Message struct {
	/*
		messages passed between front-end and back-end
	*/
	SessionID, HttpMethod string
	Body                  []datamodels.Item
}

type AckMessage struct{ Status int }

/*
	acknowledgement message
*/

func detectFailure(hostName string) {
	/*
		function to detect backend failure
		will exit after five failures
	*/
	fmt.Println("Ping, ack")
	failureCount := 0

	go func() {
		for {
			ping := Message{
				SessionID:  "1",
				HttpMethod: "PING",
				Body:       []datamodels.Item{}}

			conn, err := net.Dial("tcp", hostName)
			if err != nil {
				log.Fatal("Connection error", err)
			}
			defer conn.Close()

			ack := &AckMessage{}

			encoder := gob.NewEncoder(conn)
			encoder.Encode(ping)

			decoder := gob.NewDecoder(conn)
			decoder.Decode(ack)

			if ack.Status != 1 && failureCount == 4 {
				log.Fatal(hostName, " encountered five errors. Cya!")
			} else if ack.Status != 1 {
				t := time.Now().UTC()
				fmt.Printf("Detected failure on %s at %s \n", hostName, t)
				failureCount++
			} else {
				fmt.Printf("%s is alive \n", hostName)
			}
			time.Sleep(10 * time.Second)
		}
	}()
}

func main() {
	args := os.Args[1:]
	portNum := utils.ParseListenPort(args)
	backendHost := utils.ParseBackendHost(args)
	fmt.Println("Connecting to backendHost:", backendHost)
	app := iris.New()
	app.StaticWeb("/", "./public")

	sess := sessions.New(sessions.Config{
		Cookie: "iris_session",
	})

	ws := websocket.New(websocket.Config{})

	readingListRouter := app.Party("/readingList")

	readingListRouter.Any("/iris-ws.js", websocket.ClientHandler())

	readingListApp := mvc.New(readingListRouter)

	readingListApp.Register(
		backendservice.NewDataService(backendHost),
		sess.Start,
		ws.Upgrade,
	)
	// register controllers
	readingListApp.Handle(new(controllers.ReadingListController))

	/*
		if we detect two failures within twenty seconds kill the system
	*/
	detectFailure(backendHost)

	// start the web server
	app.Run(iris.Addr(":" + portNum))

}
