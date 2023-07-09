package execution

type Status int

const (
	ExecutionPending Status = 1
	ExecutionRunning Status = 2
	ExecutionInterupted Status = 3
	ExecutionSuccess Status = 4
	ExecutionFailed Status = 5
)