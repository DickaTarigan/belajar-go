package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Goroutine ke-", i)
		}()
	}

	wg.Wait()
	fmt.Println("Semua goroutine selesai")
}
