package main

import (
	"fmt"
	"os"

	"github.com/SamW-hash/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	newState := state{cfg: &cfg}
	newCommands := commands{cmdlist: make(map[string]func(*state, command) error)}
	newCommands.register("login", handlerLogin)
	if len(os.Args) < 2 {
		fmt.Println("Usage: gator <command> [args...]")
		os.Exit(1)
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	newCommand := command{name: cmdName, args: cmdArgs}
	if err := newCommands.run(&newState, newCommand); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
