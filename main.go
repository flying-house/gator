package main

import (
	"fmt"

	"github.com/flying-house/gator/internal/config"
)

func main() {
	user := "newUser"
	cfg, _ := config.Read()

	fmt.Println("Read 1:")
	fmt.Println(cfg)

	cfg.SetUser(user)
	cfg, _ = config.Read()

	fmt.Println(fmt.Sprintf("Read after setting user: %s\n", user))
	fmt.Println(cfg)

	return
}
