package tasks

import (
	"Phylogeny/database/queries"
	"Phylogeny/entities/models"
	"log"
	"sync"
	"time"
)

type JobQueue struct {
	jobs chan *models.Job
	wg   sync.WaitGroup
}

func NewJobQueue(queueSize int) *JobQueue {
	return &JobQueue{
		jobs: make(chan *models.Job, queueSize),
	}
}

func (jq *JobQueue) StartWorkers(workerCount int) {
	for i := 0; i < workerCount; i++ {
		go jq.worker()
	}
}

func (jq *JobQueue) worker() {
	for job := range jq.jobs {
		processJob(job)
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

func processJob(job *models.Job) {
	_ = queries.UpdateJobStatus(job, models.JobProcessing)
	log.Printf("Processing job with ID: %s", job.ID)

	time.Sleep(2 * time.Second)

	_ = queries.UpdateJobStatus(job, models.JobCompleted)
	log.Println("Completed processing job with ID:", job.ID)
}
