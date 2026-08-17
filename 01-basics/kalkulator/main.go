package main

import "fmt"
import "errors"

func tambah(x, y int) int {
	return x + y
}

func kurang(x, y int) int {
	return x - y
}

func kali(x, y int) int {
	return x * y
}

func bagi(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("tidak bisa di bagi 0")
	}
	return x / y, nil
}

func main() {
	x := 3
	y := 4
	fmt.Println(tambah(x, y))
	fmt.Println(kurang(x, y))
	fmt.Println(kali(x, y))

	hasil, err := bagi(x, y)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println(hasil)
}
