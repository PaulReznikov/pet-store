package main

import (
	"fmt"
	"github.com/PaulReznikov/pet-store/internal/config"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config: ", err)
		return
	}
}
