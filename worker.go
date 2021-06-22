package concurrency

import (
	"context"
	"log"
)

type Worker struct {
	cxt context.Context
	// 所属于的工作池
	pool chan chan Job
	// 自身的Job队列
	jobChan chan Job
}

func (w *Worker) Start() {
	go func() {
		for {
			w.pool <- w.jobChan
			select {
			case job := <-w.jobChan:
				if err := job.Do(); err != nil {
					log.Println("Do job err: ", err)
				}
			case <-w.cxt.Done():
				return
			}
		}
	}()
}

func NewWorker(ctx context.Context, pool chan chan Job) *Worker {
	return &Worker{
		cxt:     ctx,
		pool:    pool,
		jobChan: make(chan Job, 0),
	}
}
