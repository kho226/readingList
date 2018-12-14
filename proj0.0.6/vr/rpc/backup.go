package rpc

type BackupService interface {
	Prepare(args *PrepareArgs, resp *PrepareOk) error
}

type PrepareArgs struct {
	/*
		Prepare is the input argument type to Echo.
	*/
	ViewNum   int
	Request   Request
	OpNum     int
	CommitNum int
}

type PrepareOk struct {
	/*
		PrepareOk is the output type of Prepare.
	*/
	ViewNum int
	OpNum   int
	Id      int
}

type Commit struct {
	/*
		Commit is sent by primary if no new Prepare message is being sent
	*/
	ViewNum   int
	CommitNum int
}
