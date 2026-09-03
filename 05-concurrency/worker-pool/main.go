package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, hasil chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d ngerjain job %d\n", id, job)
		hasil <- job * 2 // simulasi kerjaan: kalikan 2
	}
}

func main() {
	jobs := make(chan int, 10)
	hasil := make(chan int, 10)
	var wg sync.WaitGroup

	// buat 3 worker
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, hasil, &wg)
	}

	// kirim 5 job
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs) // penting! biar worker tau gak ada job lagi

	wg.Wait()
	close(hasil)

	for r := range hasil {
		fmt.Println("Hasil:", r)
	}
}
