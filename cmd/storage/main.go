package main

import (
	"fmt"
	"log"

	"github.com/serg14159/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("Hello"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it uploaded", file)
}
