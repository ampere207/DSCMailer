package mailer

import (
	"log"
	"sync"
	"time"
)

type Job struct {
	Name            string
	Email           string
	CertificatePath string
}

func Dispatch(jobs []Job) {

	const workerCount = 3
	const rateLimit = time.Second * 2 // 1 email / 2 sec

	jobChan := make(chan Job)
	wg := sync.WaitGroup{}

	// Start workers
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(jobChan, &wg, rateLimit)
	}

	// Send jobs
	for _, job := range jobs {
		jobChan <- job
	}
	close(jobChan)

	wg.Wait()
	log.Println("All emails processed")
}

func worker(jobs <-chan Job, wg *sync.WaitGroup, rate time.Duration) {
	defer wg.Done()

	for job := range jobs {
		subject, body, err := RenderEmail(job.Name)
		if err != nil {
			log.Println("Template error:", err)
			continue
		}

		err = SendSMTP(
			job.Email,
			subject,
			body,
			job.CertificatePath,
		)

		if err != nil {
			log.Printf("Failed to send to %s: %v\n", job.Email, err)
		} else {
			log.Printf("Sent certificate to %s\n", job.Email)
		}

		time.Sleep(rate)
	}
}
