/* 
	Author: Kyle Ong
	Date: 10/25/2018

	server for front end of readinglist application
*/
package main

import (
	//standard

	"fmt"
	"os" // user-defined

	"distsys/proj0.0.4/frontend/services"
	"distsys/proj0.0.4/frontend/utils"
	"distsys/proj0.0.4/frontend/web/controllers" // third party
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/websocket"
)

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

	// start the web server
	app.Run(iris.Addr(":" + portNum))
}
