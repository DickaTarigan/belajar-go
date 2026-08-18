package main

import (
	"fmt"
	"strings"
)

func counter(kalimat string) map[string]int {
	daftarKata := strings.Fields(kalimat)

	frekuensi := make(map[string]int)

	for _, kata := range daftarKata {
		frekuensi[kata]++
	}

	return frekuensi
}

func main() {
	kalimat := strings.ToLower("Aku Anak Indonesia sehat dan kuat dan antinarkoba")
	hitung := counter(kalimat)

	fmt.Println(hitung)

	for kata, jumlah := range hitung {
		fmt.Printf("Kata: %s : %d \n", kata, jumlah)
	}
}
