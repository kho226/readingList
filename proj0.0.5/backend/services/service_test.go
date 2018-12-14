/*

 */
package service

import (
	// standard lib
	"fmt"
	"testing" // user defined

	"distsys/proj0.0.5/backend/datamodels"
	"distsys/proj0.0.5/backend/datasource"
)

func TestNewDataService(t *testing.T) {
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
