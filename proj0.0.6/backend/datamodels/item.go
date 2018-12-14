/* 
	Author: Kyle Ong
	Date: 10/25/2018

	datamodels for readinglist application
*/
package datamodels

type Item struct {
	SessionID string `json:"-"`
	ID        int64  `json:"id,omitempty"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Completed bool   `json:"completed"`
}
