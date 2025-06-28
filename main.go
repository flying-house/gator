package main

import (
	"fmt"

	"github.com/flying-house/gator/internal/config"
)

func main(error) {
	cfg, err := config.Read()
	if err != nil {
		return err
	}
	fmt.Printf(cfg.dbURL)
	return
}
