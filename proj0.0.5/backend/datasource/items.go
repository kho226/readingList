/*
	Author: Kyle Ong
	Date: 10/25/2018

	datasource for readlinglist application
*/
package datasource

import (
	"distsys/proj0.0.5/backend/datamodels"
)

var testItems = []datamodels.Item{
	datamodels.Item{
		SessionID: "test",
		Title:     "The Man in the High Castle",
		Author:    "Philip K Dick",
		Completed: false},
	datamodels.Item{
		SessionID: "test",
		Title:     "Inquisitorial Inquiries",
		Author:    "Richard L. Kagan & Abigail Dyer",
		Completed: false},
	datamodels.Item{
		SessionID: "test",
		Title:     "The Price",
		Author:    "Arthur Miller",
		Completed: false},
	datamodels.Item{
		SessionID: "test",
		Title:     "A Thousand Splendid Suns",
		Author:    "Khaled Hoesseini",
		Completed: false}}

var Items = map[string][]datamodels.Item{
	"test": testItems}
