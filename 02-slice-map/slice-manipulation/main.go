package main

import "fmt"

func reverse(angka []int) []int {
	hasil := make([]int, 0, len(angka))
	for i := len(angka) - 1; i >= 0; i-- {
		hasil = append(hasil, angka[i])
	}
	return hasil
}

func minmax(angka []int) (int, int) {
	min := angka[0]
	max := angka[0]

	for _, val := range angka {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return min, max
}

func hapusDuplikat(angka []int) []int {
	penampungan := map[int]bool{}
	var slice []int

	for _, val := range angka {
		if !penampungan[val] {
			slice = append(slice, val)
			penampungan[val] = true
		}
	}
	return slice
}

func main() {
	angka := []int{1, 3, 5, 7, 5, 3, 632, 4}

	fmt.Println("reverse:", reverse(angka))

	nilaiMin, nilaiMax := minmax(angka)
	fmt.Printf("min: %d, max: %d \n", nilaiMin, nilaiMax)

	fmt.Println("Hapus duplikat:", hapusDuplikat(angka))

}
