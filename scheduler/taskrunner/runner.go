package taskrunner

// 生产-消费者模型

type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       dataChan
	dataSize   int
	longlived  bool // 是否是长期存活的Runner
	Dispatcher fn
	Executor   fn
}

func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1), //non-block chan
		Error:      make(chan string, 1),
		Data:       make(chan interface{}, size),
		longlived:  longlived,
		dataSize:   size,
		Dispatcher: d,
		Executor:   e,
	}
}

func (r *Runner) startDispatch() {
	defer func() {
		if !r.longlived {
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()

	// 异步接收chan数据
	for {
		select { // select is non-blocking
		case c := <-r.Controller:
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXECUTE
				}
			} else if c == READY_TO_EXECUTE {
				err := r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e := <-r.Error:
			if e == CLOSE {
				return // exit
			}
		default:

		}
	}

}

func (r *Runner) StartAll() {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}
