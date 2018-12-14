package rpc

// ViewService is the RPC to perform a view change.
type ViewService interface {
	// StartViewChange initiates a view change.
	StartViewChange(*StartViewChangeArgs, *StartViewChangeResp) error
	// DoViewChange tells the new primary to start a new view.
	DoViewChange(*DoViewChangeArgs, *DoViewChangeResp) error
	// StartView tells the backups to transition to a new view.
	StartView(*StartViewArgs, *StartViewResp) error
}

type StartViewChangeArgs struct {
	/*
		// StartViewChangeArgs is the arguments to start a view change.
	*/
	ViewNum int
	Id      int
}

type StartViewChangeResp struct {
	/*
		StartViewChangeResp is the response to a StartViewChange message.
	*/
}

type DoViewChangeArgs struct {
	/*
		DoViewChangeArgs is the arguments to tell the new primary to start a new view.
	*/
	ViewNum             int
	Log                 []OpRequest
	LatestNormalViewNum int
	OpNum               int
	CommitNum           int
	Id                  int
}

type DoViewChangeResp struct {
	/*
		DoViewChangeResp is the response to a DoViewChange message.
	*/
}

type StartViewArgs struct {
	/*
		StartViewArgs is the arguments for the primary to start a new view.
	*/
	ViewNum   int
	Log       []OpRequest
	OpNum     int
	CommitNum int
}

type StartViewResp struct {
	/*
		StartViewResp is the response to a StartView message.
	*/
}
