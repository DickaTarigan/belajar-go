package main

import (
	"fmt"
	"math"
)

type BangunDatar interface {
	Luas() float64
}

type Lingkaran struct {
	r float64
}

func (l Lingkaran) Luas() float64 {
	return math.Pi * l.r * l.r
}

type Persegi struct {
	s float64
}

func (p Persegi) Luas() float64 {
	return p.s * p.s
}

type Segitiga struct {
	a float64
	t float64
}

func (s Segitiga) Luas() float64 {
	return s.a * s.t / 2
}

func HitungLuas(b BangunDatar) float64 {
	return b.Luas()
}

func main() {
	lingkaran := Lingkaran{r: 10}
	persegi := Persegi{s: 4}
	segitiga := Segitiga{
		a: 5,
		t: 6,
	}

	fmt.Println("Luas lingkaran:", HitungLuas(lingkaran))
	fmt.Println("Luas persegi:", HitungLuas(persegi))
	fmt.Println("Luas segitiga:", HitungLuas(segitiga))
}
