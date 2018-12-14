package rpc

type RecoveryService interface {
	/*
		RecoveryService is the RPC to perform a recovery.
	*/
	Recover(request *RecoveryRequest, response *RecoveryResponse) error
}

type RecoveryRequest struct {
	/*
		RecoveryRequest is the request to start a recovery.
	*/
	Id    int
	Nonce int
}

type RecoveryResponse struct {
	/*
		RecoveryRequest is the response to a recovery request.
	*/
	ViewNum   int
	Nonce     int
	Log       []OpRequest
	OpNum     int
	CommitNum int
	Id        int
	Mode      string
}
