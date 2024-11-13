package tasks

import (
	"Phylogeny/database/queries"
	"Phylogeny/entities/models"
	"log"
	"sync"
	"time"
)

type BroadcastChannel struct {
	clients map[string][]chan string
	mu      sync.RWMutex
}

func NewBroadcastChannel() *BroadcastChannel {
	return &BroadcastChannel{
		clients: make(map[string][]chan string),
	}
}

func (bc *BroadcastChannel) Subscribe(jobID string) chan string {
	ch := make(chan string, 1)
	bc.mu.Lock()
	bc.clients[jobID] = append(bc.clients[jobID], ch)
	bc.mu.Unlock()
	return ch
}

func (bc *BroadcastChannel) Unsubscribe(jobID string, ch chan string) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	for i, client := range bc.clients[jobID] {
		if client == ch {
			bc.clients[jobID] = append(bc.clients[jobID][:i], bc.clients[jobID][i+1:]...)
			close(ch)
			break
		}
	}
}

func (bc *BroadcastChannel) Notify(jobID, status string) {
	bc.mu.RLock()
	defer bc.mu.RUnlock()
	for _, ch := range bc.clients[jobID] {
		ch <- status
	}
}

type JobQueue struct {
	jobs      chan *models.Job
	wg        sync.WaitGroup
	Broadcast *BroadcastChannel
}

func NewJobQueue(queueSize int) *JobQueue {
	return &JobQueue{
		jobs:      make(chan *models.Job, queueSize),
		Broadcast: NewBroadcastChannel(),
	}
}

func (jq *JobQueue) StartWorkers(workerCount int) {
	for i := 0; i < workerCount; i++ {
		go jq.worker()
	}
}

func (jq *JobQueue) worker() {
	for job := range jq.jobs {
		time.Sleep(1 * time.Second)
		jq.Broadcast.Notify(job.ID.String(), string(models.JobQueued))
		log.Printf("Queued job with ID: %s", job.ID)
		time.Sleep(4 * time.Second)
		processJob(jq, job)
		jq.wg.Done()
	}
}

func (jq *JobQueue) Enqueue(job *models.Job) {
	jq.wg.Add(1)
	jq.jobs <- job
}

func (jq *JobQueue) Shutdown() {
	log.Println("Shutting down job queue...")

	done := make(chan struct{})
	go func() {
		jq.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		log.Println("All jobs processed, shutting down workers.")
	case <-time.After(10 * time.Second):
		log.Println("Timeout reached. Forcing shutdown.")
	}

	close(jq.jobs)
}

func processJob(jq *JobQueue, job *models.Job) {
	_ = queries.UpdateJobStatus(job, models.JobProcessing)
	jq.Broadcast.Notify(job.ID.String(), string(models.JobProcessing))
	log.Printf("Processing job with ID: %s", job.ID)

	time.Sleep(2 * time.Second)

	_ = queries.UpdateJobStatus(job, models.JobCompleted)
	jq.Broadcast.Notify(job.ID.String(), string(models.JobCompleted))
	log.Println("Completed processing job with ID:", job.ID)
}
