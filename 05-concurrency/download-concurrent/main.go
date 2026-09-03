package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Result struct {
	url      string
	bytes    int64
	duration time.Duration
	err      error
}

// Fungsi pembantu untuk fetch satu URL
func fetch(url string) Result {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return Result{url: url, err: err, duration: time.Since(start)}
	}
	defer resp.Body.Close()

	// Baca seluruh response body untuk simulasi download riil
	n, err := io.Copy(io.Discard, resp.Body)
	return Result{
		url:      url,
		bytes:    n,
		duration: time.Since(start),
		err:      err,
	}
}

func downloadSequential(urls []string) {
	start := time.Now()
	fmt.Println("=== 1. Menjalankan Sekuensial ===")

	for _, u := range urls {
		res := fetch(u)
		if res.err != nil {
			fmt.Printf("Gagal %s: %v\n", res.url, res.err)
			continue
		}
		fmt.Printf("Selesai %-20s (%6d bytes) dalam %v\n", res.url, res.bytes, res.duration)
	}

	fmt.Printf("Total durasi sekuensial: %v\n\n", time.Since(start))
}

func downloadConcurrent(urls []string) {
	start := time.Now()
	fmt.Println("=== 2. Menjalankan Konkuren (Goroutine + Channel) ===")

	// Gunakan buffered channel seukuran jumlah URL agar pengirim tidak terblokir
	ch := make(chan Result, len(urls))

	for _, u := range urls {
		go func(url string) {
			ch <- fetch(url)
		}(u)
	}

	// Panen hasil dari channel sebanyak jumlah URL yang dikirim
	for i := 0; i < len(urls); i++ {
		res := <-ch
		if res.err != nil {
			fmt.Printf("Gagal %s: %v\n", res.url, res.err)
			continue
		}
		fmt.Printf("Selesai %-20s (%6d bytes) dalam %v\n", res.url, res.bytes, res.duration)
	}

	fmt.Printf("Total durasi konkuren: %v\n", time.Since(start))
}

func main() {
	urls := []string{
		"https://go.dev",
		"https://pkg.go.dev",
		"https://github.com",
	}

	downloadSequential(urls)
	downloadConcurrent(urls)
}
