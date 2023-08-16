package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents the task to be executed by a worker
type Job struct {
	ID int
}

// WorkerPool represents a pool of worker goroutines
type WorkerPool struct {
	numWorkers int
	jobQueue   chan Job
	results    chan int
	wg         sync.WaitGroup
}

// NewWorkerPool creates a new worker pool with the specified number of workers
func NewWorkerPool(numWorkers, jobQueueSize int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		jobQueue:   make(chan Job, jobQueueSize),
		results:    make(chan int, jobQueueSize),
	}
}

// worker function to process jobs from the queue
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for job := range wp.jobQueue {
		// Do the actual work here
		fmt.Printf("Worker %d started job %d\n", id, job.ID)
		time.Sleep(time.Second) // Simulating work
		fmt.Printf("Worker %d finished job %d\n", id, job.ID)
		wp.results <- job.ID
	}
}

// Start starts the worker pool and dispatches jobs to workers
func (wp *WorkerPool) Start() {
	for i := 1; i <= wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// Wait waits for all workers to finish and closes the results channel
func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
	close(wp.results)
}

// AddJob adds a job to the job queue
func (wp *WorkerPool) AddJob(job Job) {
	wp.jobQueue <- job
}

// CollectResults collects and prints results from the results channel
func (wp *WorkerPool) CollectResults() {
	for result := range wp.results {
		fmt.Printf("Result received for job %d\n", result)
	}
}

func main() {
	numWorkers := 3
	numJobs := 10

	workerPool := NewWorkerPool(numWorkers, numJobs)

	// Adding jobs to the Job Queue
	for i := 1; i <= numJobs; i++ {
		workerPool.AddJob(Job{ID: i})
	}

	close(workerPool.jobQueue)

	workerPool.Start()
	workerPool.Wait()
	workerPool.CollectResults()
}
