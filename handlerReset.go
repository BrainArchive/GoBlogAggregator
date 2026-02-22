package main

import (
	"context"
	"fmt"
	"log"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		log.Fatalf("error deleting users: %v \n", err)
	}

	err = s.ConfigPtr.SetUser("")
	if err != nil {
		log.Fatalf("error resetting json config: %v \n", err)
	}
	fmt.Println("users table is reset")
	return nil
}
