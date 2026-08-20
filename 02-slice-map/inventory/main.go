package main

import (
	"fmt"
	"strings"
)

type Item struct {
	Nama  string
	Harga int
}

func tambah(inv []Item, nama string, harga int) []Item {
	itemBaru := Item{
		Nama:  nama,
		Harga: harga,
	}

	inv = append(inv, itemBaru)

	return inv

}

func cari(inv []Item, nama string) (Item, bool) {
	for _, item := range inv {
		if strings.EqualFold(item.Nama, nama) {
			return item, true
		}
	}

	return Item{}, false
}

func hapus(inv []Item, nama string) []Item {
	var keranjang []Item

	for _, item := range inv {
		if !strings.EqualFold(item.Nama, nama) {
			keranjang = append(keranjang, item)
		}
	}
	return keranjang
}

func main() {
	var inventory []Item

	inventory = tambah(inventory, "Laptop", 15000)
	inventory = tambah(inventory, "Smartphone", 2000)

	fmt.Println(inventory)

	barang, ditemukan := cari(inventory, "Keyboard")
	if ditemukan {
		fmt.Printf("Ditemukan: Barang: %s, Harga: %d \n", barang.Nama, barang.Harga)
	} else {
		fmt.Println("Tidak ditemukan")
	}

	inventory = hapus(inventory, "Laptop")
	fmt.Println(inventory)

}
