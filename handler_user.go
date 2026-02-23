package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SamW-hash/gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("You must provide a username.")
	}
	name := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		if err == sql.ErrNoRows {
			os.Exit(1)
		}
		return err
	}
	if err := s.cfg.SetUser(cmd.args[0]); err != nil {
		return err
	}
	fmt.Println("User has been successfully set.")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("You must provide a username.")
	}
	name := cmd.args[0]
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}
	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				os.Exit(1)
			}
		}
		return err
	}
	log.Printf("created user: %+v\n", user)
	return s.cfg.SetUser(name)
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}
	if err := s.db.Wipe(context.Background()); err != nil {
		return fmt.Errorf("Reset unsuccessful: %w", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}
