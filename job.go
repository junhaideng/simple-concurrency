package concurrency

type Job interface{
	Do() error 
}