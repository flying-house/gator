package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/flying-house/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmdReg map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Bad input - try 'login [username]'")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("Set user:", cmd.args[0])
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	return c.cmdReg[cmd.name](s, cmd)
}

func (c *commands) register(name string, fn func(*state, command) error) {
	c.cmdReg[name] = fn
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Failed to read config file")
	}
	s := state{&cfg}

	cmds := commands{
		cmdReg: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %v <command> [args...]", os.Args[0])
	}
	cmd := command{name: os.Args[1], args: os.Args[2:]}
	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatalf("Failed to execute '%v': %v", cmd.name, err)
	}
}
