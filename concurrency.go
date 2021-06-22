package concurrency

import (
	"context"
)

type Dispatcher struct {
	ctx context.Context
	// 所管理的工作池
	pool chan chan Job
	// 管理的所有worker
	workers  []*Worker
	JobQueue chan Job
}

// NewDispatcher 返回一个声明工作池大小和工作队列长度的分发器
func NewDispatcher(ctx context.Context, poolSize int, jobSize int) *Dispatcher {
	pool := make(chan chan Job, poolSize)
	workers := make([]*Worker, poolSize)
	for i := 0; i < poolSize; i++ {
		worker := NewWorker(ctx, pool)
		workers[i] = worker
	}
	return &Dispatcher{
		ctx:      ctx,
		pool:     pool,
		workers:  workers,
		JobQueue: make(chan Job, jobSize),
	}
}

func (d *Dispatcher) Run() {
	for _, worker := range d.workers {
		worker.Start()
	}
	go func() {
		for {
			select {
			case job := <-d.JobQueue:
				// 挑选出一个工作队列
				jobChan := <-d.pool
				// 添加Job
				jobChan <- job
			case <-d.ctx.Done():
				return
			}
		}
	}()
}
