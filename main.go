package main

import (
	"fmt"
	"log"

	"github.com/bootdotdev/boot-dev-gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("erro reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
