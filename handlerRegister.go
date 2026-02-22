package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/brainarchive/goblogaggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("usage gator register <username>")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err == nil {
		fmt.Printf("%s already exists!\n", cmd.args[0])
		os.Exit(1)
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})

	if err != nil {
		fmt.Printf("Error when creating User: %v\n", err)
		os.Exit(1)
	}

	err = s.ConfigPtr.SetUser(cmd.args[0])
	if err != nil {
		fmt.Printf("Error when creating User: %v\n", err)
		return err
	}

	fmt.Println("User is Registered")
	s.ConfigPtr.CurrentUserName = user.Name
	return nil
}
