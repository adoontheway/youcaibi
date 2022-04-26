package taskrunner

const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE  = "e"
	CLOSE             = "c"
)

//
type controlChan chan string

//
type dataChan chan interface{}

// dispatcher & executor
type fn func(dc dataChan) error
