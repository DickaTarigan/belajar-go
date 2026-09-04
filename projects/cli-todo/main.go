package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	ID      int
	Judul   string
	Selesai bool
}

func saveTodos(namaFile string, data []Todo) error {
	jsonData, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		return fmt.Errorf("Gagal convert ke JSON %w", err)
	}

	err = os.WriteFile(namaFile, jsonData, 0664)
	if err != nil {
		return fmt.Errorf("Gagal menyimpan file %w", err)
	}
	return nil
}

func loadTodos(namaFile string) ([]Todo, error) {
	Hasil := []Todo{}
	jsonData, err := os.ReadFile(namaFile)
	if err != nil {
		//Kembalikan struct kosong jika file blum dibuat
		if errors.Is(err, os.ErrNotExist) {
			return []Todo{}, nil
		}

		return nil, fmt.Errorf("Gagal membaca data %w", err)
	}

	//Jika file ada tapi data tidak ada
	if len(jsonData) == 0 {
		return []Todo{}, nil
	}

	err = json.Unmarshal(jsonData, &Hasil)
	if err != nil {
		return nil, fmt.Errorf("Gagal convert ke struct")
	}
	return Hasil, nil
}

func getNextID(todos []Todo) int {
	maxID := 0
	for _, v := range todos {
		if v.ID > maxID {
			maxID = v.ID
		}
	}
	return maxID + 1
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: main.go <command> [args]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Judul ToDo tidak boleh kosong")
			fmt.Println("Usage: main.go add \"Judul ToDo\"")
			return
		}
		judul := os.Args[2]

		todos, err := loadTodos("todos.json")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		todoBaru := Todo{
			ID:      getNextID(todos),
			Judul:   judul,
			Selesai: false,
		}

		todos = append(todos, todoBaru)

		//Simpan menjadi json
		err = saveTodos("todos.json", todos)
		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Println("Todo Berhasil Ditambahkan:", judul)

	case "list":
		fmt.Println("Mau liat todo")
	case "done":
		fmt.Println("Mau tandain selesai")
	case "delete":
		fmt.Println("Mau Hapus todo")
	default:
		fmt.Println("Command tidak dikenal:", command)
	}
}
