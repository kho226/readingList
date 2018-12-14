// rpc defines all the RPC interfaces.
package rpc

import (
	"distsys/proj0.0.6/datamodels"
)

type VrgoService interface {
	/*
		VrgoService defines the APIs Vrgo exposes to users.
	*/
	Execute(*Request, *Response) error
}

// Request is the input argument type to RequestRPC.
type Request struct {
	Op         Operation
	ClientId   int
	RequestNum int
	// Do we need view number as well?
}

type OpRequest struct {
	/*
		OpRequest represents an operation record that has a Request and a operation number.
	*/
	Request Request
	OpNum   int
}

type Response struct {
	/*
		Response is the output type of RequestRPC.
	*/
	ViewNum    int
	RequestNum int
	OpResult   OperationResult
	Err        string
}

type Message struct {
	/*
		messages passed between front-end and back-end
	*/
	SessionID, HttpMethod string
	Body                  []datamodels.Item
}

type Operation struct {
	/*
		Operation is the user operation.
	*/
	Message Message
}

type OperationResult struct {
	/*
		OperationResult is the result of the user operation.
	*/
	Message Message
}
