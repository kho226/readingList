// oplog provides the interface to the in-memory log.
package oplog

import (
	"context"
	"fmt"
	"log"

	"distsys/proj0.0.6/vr/rpc"
)

// OpRequestLog is the in-memory log to store all the records.
type OpRequestLog struct {
	Requests []rpc.OpRequest
}

// New creates an OpRequestLog.
func New() *OpRequestLog {
	return &OpRequestLog{}
}

func (o *OpRequestLog) AppendRequest(ctx context.Context, request *rpc.Request, opNum int) error {
	/*
		AppendRequest appends a request along with its opNum to the log.
	*/
	log.Printf("oplog adding %v at opNum %v", request, opNum)
	r := rpc.OpRequest{Request: *request, OpNum: opNum}
	o.Requests = append(o.Requests, r)
	return nil
}

func (o *OpRequestLog) ReadLast(ctx context.Context) (*rpc.Request, int, error) {
	/*
		ReadLast returns the last request from the log or an error if the log is empty.
	*/
	if len(o.Requests) == 0 {
		return nil, 0, fmt.Errorf("OpRequestLog is empty")
	}

	r := o.Requests[len(o.Requests)-1]

	return &r.Request, r.OpNum, nil
}

func (o *OpRequestLog) Undo(ctx context.Context) {
	/*
		Undo removes the last record from the log.
	*/
	o.Requests = o.Requests[:len(o.Requests)-1]
}
