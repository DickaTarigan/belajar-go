package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
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
		if len(os.Args) != 3 {
			fmt.Println("Error: Judul Harus diapit tanda kutip jika lebih dari 1 kata")
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
			return
		}

		fmt.Println("Todo Berhasil Ditambahkan:", judul)

	case "list":
		todos, err := loadTodos("todos.json")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(todos) == 0 {
			fmt.Println("Belum ada data Todo")
			return
		}

		for _, todo := range todos {
			if todo.Selesai {
				fmt.Printf("%d. [X] %s \n", todo.ID, todo.Judul)
			} else {
				fmt.Printf("%d. [ ] %s \n", todo.ID, todo.Judul)
			}
		}

	case "done":
		if len(os.Args) != 3 {
			fmt.Println("Error: Harus menggunakan 1 ID")
			fmt.Println("Usage: main.go done <ID>")
			return
		}

		//Konversi string di cli menjadi int
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: ID harus bilangan bulat", err)
			return
		}

		todos, err := loadTodos("todos.json")
		if err != nil {
			fmt.Print("Error:", err)
			return
		}

		//cari id
		ditemukan := false

		for i := range todos {
			if id == todos[i].ID {
				todos[i].Selesai = true
				ditemukan = true
				break
			}
		}

		//Jika tidak ditemukan
		if !ditemukan {
			fmt.Println("Error: ID tidak ditemukan")
			return
		}

		err = saveTodos("todos.json", todos)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("ID %d Berhasil di tandai selesai \n", id)

	case "delete":
		if len(os.Args) != 3 {
			fmt.Println("Error: Harus menggunakan 1 ID")
			fmt.Println("Usage: main.go delete <ID>")
			return
		}

		//Konversi string di cli menjadi int
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: ID harus bilangan bulat", err)
			return
		}

		todos, err := loadTodos("todos.json")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		//cari id
		ditemukan := false

		for i := range todos {
			if id == todos[i].ID {
				todos = append(todos[:i], todos[i+1:]...)
				ditemukan = true
				break
			}
		}

		//Jika tidak ditemukan
		if !ditemukan {
			fmt.Println("Error: ID tidak ditemukan")
			return
		}

		err = saveTodos("todos.json", todos)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("ID %d Berhasil di hapus \n", id)

	default:
		fmt.Println("Command tidak dikenal:", command)
	}
}
