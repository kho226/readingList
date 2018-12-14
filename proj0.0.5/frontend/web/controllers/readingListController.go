/*
	Author: Kyle Ong
	Date: 10/25/2018

	viewcontroller for readinglist app
*/
package controllers

import (
	"fmt"

	"distsys/proj0.0.5/frontend/datamodels"
	"distsys/proj0.0.5/frontend/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/websocket"
)

type ReadingListController struct {
	BackendService backendservice.Service

	Session *sessions.Session
}

func (c *ReadingListController) BeforeActivation(b mvc.BeforeActivation) {
	b.Dependencies().Add(func(ctx iris.Context) (items []datamodels.Item) {
		ctx.ReadJSON(&items)
		return
	})
}

func (c *ReadingListController) Get() []datamodels.Item {
	fmt.Println("ReadingListController.get()", c.Session.ID())
	payload := c.BackendService.Get(c.Session.ID())
	return payload
}

type PostItemResponse struct {
	Success bool `json:"success"`
}

var emptyResponse = PostItemResponse{Success: false}

func (c *ReadingListController) Post(newItems []datamodels.Item) PostItemResponse {
	if err := c.BackendService.Save(c.Session.ID(), newItems); err != nil {
		return emptyResponse
	}

	return PostItemResponse{Success: true}
}

func (c *ReadingListController) GetSync(conn websocket.Connection) {
	conn.Join(c.Session.ID())
	conn.On("save", func() {
		conn.To(c.Session.ID()).Emit("saved", nil)
	})

	conn.Wait()
}
