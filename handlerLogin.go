package main

import (
	"context"
	"errors"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("usage gator login <username>")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		fmt.Printf("%s does not exist.\n", cmd.args[0])
		os.Exit(1)
	}

	err = s.ConfigPtr.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("User is Set")
	return nil
}
