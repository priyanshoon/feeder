package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/priyanshoon/feeder/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>\n", cmd.Name)
	}

	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("user name does not exist: %w\n", err)
	}

	err = s.config.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w\n", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>\n", cmd.Name)
	}

	name := cmd.Args[0]

	user := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	saveUser, err := s.db.CreateUser(context.Background(), user)
	if err != nil {
		return fmt.Errorf("cannot save the user to database: %w\n", err)
	}

	err = s.config.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set the current user %w\n", err)
	}

	fmt.Println("users has been created...")

	log.Printf("id: %v, name: %s, created_at: %v, updated_at: %v\n", saveUser.ID, saveUser.Name,
		saveUser.CreatedAt, saveUser.UpdatedAt)

	return nil
}
