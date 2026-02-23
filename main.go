package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/SamW-hash/gator/internal/config"
	"github.com/SamW-hash/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	dbQueries := database.New(db)
	newState := state{db: dbQueries, cfg: &cfg}
	newCommands := commands{cmdlist: make(map[string]func(*state, command) error)}
	newCommands.register("login", handlerLogin)
	if len(os.Args) < 2 {
		fmt.Println("Usage: gator <command> [args...]")
		os.Exit(1)
	}
	newCommands.register("register", handlerRegister)
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	newCommand := command{name: cmdName, args: cmdArgs}
	if err := newCommands.run(&newState, newCommand); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
