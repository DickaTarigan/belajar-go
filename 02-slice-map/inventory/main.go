package main

import "fmt"

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

func cari(inv []Item, nama string) bool {
	itemCari := Item{
		Nama: nama,
	}

	for _, item := range inv {
		if item.Nama == itemCari.Nama {
			return true
		}
	}

	return false
}

func hapus(inv []Item, nama string) []Item {
	var keranjang []Item

	itemHapus := Item{
		Nama: nama,
	}

	for _, item := range inv {
		if item.Nama != itemHapus.Nama {
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

	fmt.Println(cari(inventory, "Meki"))
	fmt.Println(cari(inventory, "Laptop"))

	fmt.Println(hapus(inventory, "Ayam"))
	fmt.Println(hapus(inventory, "Laptop"))

}
