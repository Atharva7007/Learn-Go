package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"sync"
	"time"
)

type LogJob struct {
	ID   int
	Line string
}

type SafeCounter struct {
	err_count int
	mu        sync.Mutex
}

func worker(i int, jobs <-chan LogJob, sc *SafeCounter) {
	for job := range jobs {
		fmt.Println("Worker", i, "working on Job ID", job.ID)
		if strings.Contains(job.Line, "ERROR") {
			sc.mu.Lock()
			sc.err_count++
			sc.mu.Unlock()
		}
		time.Sleep(3 * time.Second)
	}
}

func main() {
	start := time.Now()
	jobs := make(chan LogJob)
	var wg sync.WaitGroup
	var sc SafeCounter

	err_info := map[int]string{1: "ERROR", 2: "INFO"}

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker(i, jobs, &sc)
		})
	}

	for i := 1; i <= 10; i++ {
		job := LogJob{ID: i, Line: err_info[1+rand.IntN(2)]}
		fmt.Println(job)
		jobs <- job
	}
	close(jobs)
	wg.Wait()
	end := time.Now()

	fmt.Printf("Final error count: %d", sc.err_count)
	fmt.Printf("\nTime taken: %v\n", end.Sub(start))
}
