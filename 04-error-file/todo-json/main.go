package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	ID      int
	Judul   string
	Selesai bool
}

func save(namaFile string, data []Todo) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("gagal convert ke json: %w", err)
	}

	err = os.WriteFile(namaFile, jsonData, 0664)
	if err != nil {
		return fmt.Errorf("gagal menulis file: %w", err)
	}
	return nil
}

func load(namaFile string) ([]Todo, error) {
	Hasil := []Todo{}
	jsonData, err := os.ReadFile(namaFile)
	if err != nil {
		return nil, fmt.Errorf("File gagal di load: %w", err)
	}

	err = json.Unmarshal(jsonData, &Hasil)
	if err != nil {
		return nil, fmt.Errorf("gagal mengubah menjadi struct: %w", err)
	}

	return Hasil, nil
}

func main() {
	daftarTodo := []Todo{
		{ID: 1, Judul: "Belajar error go", Selesai: true},
		{ID: 2, Judul: "Belajar convert json", Selesai: true},
	}

	err := save("Belajar convert.json", daftarTodo)
	if err != nil {
		fmt.Println("Gagal menyimpan:", err)
		return
	}
	fmt.Println("Data berhasil disimpan")

	jsonToStruct, err := load("Belajar convert.json")
	if err != nil {
		fmt.Println("Gagal convert ke struct:", err)
		return
	}
	fmt.Println(jsonToStruct)
}
