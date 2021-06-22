package main

import (
	"context"
	"fmt"
	con "github.com/junhaideng/simple-concurrency"
	"net/http"
)

var (
	dispatcher = con.NewDispatcher(context.TODO(), 100, 100)
)

type PayLoad struct {
}

// myJob implements Job interface
type myJob struct {
	PayLoad
}

func (m *myJob) Do() error {
	fmt.Printf("do job: %#v\n", m.PayLoad)
	return nil
}

func NewJob(payload PayLoad) con.Job {
	return &myJob{
		PayLoad: payload,
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	payload := PayLoad{}
	dispatcher.JobQueue <- NewJob(payload)
	w.Write([]byte("Hello world"))
}

func main() {
	dispatcher.Run()

	http.HandleFunc("/payload", handle)
	http.ListenAndServe(":8080", nil)
}
