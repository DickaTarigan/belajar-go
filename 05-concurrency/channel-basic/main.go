package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2) // buffer 2 slot

	go func() {
		fmt.Println("Goroutine: mau kirim pesan 1...")
		ch <- "pesan 1"
		fmt.Println("Goroutine: pesan 1 terkirim!") // ini bisa LANGSUNG muncul, gak nunggu penerima
		fmt.Println("Goroutine: mau kirim pesan 2...")
		ch <- "pesan 2"
		fmt.Println("Goroutine: pesan 2 terkirim!")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Main: baru sekarang mau nerima...")
	fmt.Println("Main terima:", <-ch)
	fmt.Println("Main terima:", <-ch)
}
