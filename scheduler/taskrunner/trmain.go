package taskrunner

import "time"

// timer -> setup -> start { trigger -> taskrunner}

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startworker() {
	for {
		select {
		case <-w.ticker.C:
			go w.runner.StartAll()
		}
	}
}
