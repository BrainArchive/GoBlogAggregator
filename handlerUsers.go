package main

import (
	"context"
	"fmt"
	"log"
)

func handlerUsers(s *state, cmd command) error {

	userList, err := s.db.ListUsers(context.Background())
	if err != nil {
		log.Fatalf("fatal error occurred listing users: %v\n", err)
	}
	for _, user := range userList {
		var userStr string
		if user == s.ConfigPtr.CurrentUserName {
			userStr = user + " (current)"
		} else {
			userStr = user
		}
		fmt.Println(userStr)
	}

	return nil
}
