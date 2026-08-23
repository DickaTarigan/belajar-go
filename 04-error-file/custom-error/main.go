package main

import "fmt"

type ValidasiError struct {
	Pesan string
}

func (v ValidasiError) Error() string {
	return v.Pesan
}

func ValidasiUmur(umur int) error {
	if umur <= 0 {
		return ValidasiError{Pesan: "Umur tidak boleh kurang dari 0"}
	}
	return nil
}

func main() {
	if err := ValidasiUmur(-1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Umur valid")
	}

	if err := ValidasiUmur(1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Umur valid")
	}

}
