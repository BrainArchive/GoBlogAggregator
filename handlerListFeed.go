package main

import (
	"context"
	"fmt"
)

func handlerListFeed(s *state, cmd command) error {
	feedsList, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return err
	}
	// you could probably make this concurrent.
	// split your feedsList into chunks.
	// Those chunks are then rendered
	for _, feed := range feedsList {
		fmt.Println("====================")
		fmt.Printf("%v\n", feed.Name)
		fmt.Printf("%v\n", feed.Url)
		fmt.Printf("%v\n", feed.UserID)
		user, err := s.db.GetUserFromId(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", user.Name)
		fmt.Println("====================")
	}

	return nil
}
