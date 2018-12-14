/*
	Author: Kyle Ong
	Date: 10/25/2018

	test suite for services
*/
package service

import (
	// standard lib
	"fmt"
	"testing" // user defined

	"distsys/proj0.0.6/datamodels"
	"distsys/proj0.0.6/backend/datasource"
)

func TestNewDataService(t *testing.T) {
	/*
		tests new data service
	*/
	dataService := NewDataService(datasource.Items)
	if _, exists := dataService.items["test"]; !exists {
		message := fmt.Sprintf("Data Service does not have key 'test'")
		t.Error(message)
	}
	testItems := dataService.items["test"]
	for _, item := range testItems {
		if item.SessionID != "test" {
			message := fmt.Sprintf("Expected: test. got: %s", item.SessionID)
			t.Error(message)
		}
	}
}

func TestGet(t *testing.T) {
	/*
		tests  GET
	*/
	dataService := NewDataService(datasource.Items)
	testItems := dataService.Get("test")
	for _, item := range testItems {
		if item.SessionID != "test" {
			message := fmt.Sprintf("Expected: test. got: %s", item.SessionID)
			t.Error(message)
		}
	}
}

func TestSave(t *testing.T) {
	/*
		tests POST
	*/
	dataService := NewDataService(datasource.Items)
	saveItem := datamodels.Item{
		SessionID: "saved_sessID",
		Title:     "saved_title",
		Completed: true}
	saveItems := []datamodels.Item{
		saveItem,
	}
	err := dataService.Save("test", saveItems)
	if err != nil {
		t.Error(err)
	}
}
