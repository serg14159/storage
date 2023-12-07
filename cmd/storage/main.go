package main

import (
	"fmt"

	"github.com/serg14159/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	fmt.Println("it work", st)
}
