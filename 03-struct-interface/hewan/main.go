package main

import "fmt"

type Hewan interface {
	Bersuara() string
	Bernama() string
}
type Kucing struct {
	Nama string
}

func (k *Kucing) Bersuara() string {
	return "MIAWWW"
}
func (k *Kucing) Bernama() string {
	return k.Nama
}

type Anjing struct {
	Nama string
}

func (a *Anjing) Bersuara() string {
	return "Gong Gong"
}
func (a *Anjing) Bernama() string {
	return a.Nama
}

func CetakSuara(h Hewan) {
	fmt.Printf("Nama Hewan: %s, Suara: %s\n", h.Bernama(), h.Bersuara())
}

func main() {
	kucing := Kucing{
		Nama: "Asko",
	}

	anjing := Anjing{
		Nama: "Black",
	}

	CetakSuara(&kucing)
	CetakSuara(&anjing)
}
