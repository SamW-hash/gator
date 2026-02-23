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
	if len(os.Args) < 2 {
		fmt.Println("Usage: gator <command> [args...]")
		os.Exit(1)
	}

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
	defer db.Close()

	st := state{db: database.New(db), cfg: &cfg}

	cmds := commands{cmdlist: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)

	cmd := command{name: os.Args[1], args: os.Args[2:]}
	if err := cmds.run(&st, cmd); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
