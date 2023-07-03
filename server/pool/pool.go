package pool

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Job struct {
	ID      int64
	JobFunc func(param JobParam)
	Param   JobParam
}

type JobParam struct {
	Project string
	Domains []string
}

type Worker struct {
	ID         int
	JobChannel chan Job
	Quit       chan bool
}

type Pool struct {
	JobQueue     chan Job
	WorkerNumber int
	Workers      []Worker
	wg           sync.WaitGroup
}

var ThreadPool *Pool

func RegisterPool() {
	ThreadPool = NewPool(10, 200)
	ThreadPool.Wait()
	// 注册关闭钩子，用于在关闭时调用线程池的 Shutdown 方法
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		ThreadPool.Shutdown()
		fmt.Println("Shutting down...")
		os.Exit(0)
	}()

}

func NewWorker(id int, jobChannel chan Job) Worker {
	return Worker{
		ID:         id,
		JobChannel: jobChannel,
		Quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.JobChannel:
				fmt.Printf("Worker %d started job %d\n", w.ID, job.ID)
				job.JobFunc(job.Param)
				fmt.Printf("Worker %d finished job %d\n", w.ID, job.ID)
			case <-w.Quit:
				return
			}
		}
	}()
}

func NewPool(workerNumber, jobQueueSize int) *Pool {
	jobQueue := make(chan Job, jobQueueSize)
	workers := make([]Worker, workerNumber)

	pool := &Pool{
		JobQueue:     jobQueue,
		WorkerNumber: workerNumber,
		Workers:      workers,
	}

	for i := 0; i < workerNumber; i++ {
		worker := NewWorker(i, jobQueue)
		workers[i] = worker
		worker.Start()
	}

	return pool
}

func (p *Pool) AddJob(job Job) {
	p.wg.Add(1)
	p.JobQueue <- job
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func (p *Pool) Shutdown() {
	for _, worker := range p.Workers {
		worker.Quit <- true
	}
}
