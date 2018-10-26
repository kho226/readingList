/* 
	Author: Kyle Ong
	Date: 10/25/2018

	backendservice tests for readinglist application

	todo
	- [ ] mock backend server to test GET and SAVE
*/
package backendservice

import (
	// standard lib
	"fmt"
	"testing"

)

func TestNewBackendDataService(t *testing.T) {
	dataService := NewDataService("localhost:8080")
	if dataService.HostName != "localhost:8080"{
		message := fmt.Sprintf("Expected: localhost:8080. got: %s", dataService.HostName)
		t.Error(message)

	}
}