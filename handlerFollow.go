package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/brainarchive/goblogaggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return errors.New("usage gator follow <url>")
	}

	feed, err := s.db.FeedFromUrl(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		fmt.Printf("Error Creating Feed Follow\n")
		return err
	}

	fmt.Printf("ID: %v\n", feedFollow.ID)
	fmt.Printf("Created At: %v\n", feedFollow.CreatedAt)
	fmt.Printf("Updated At: %v\n", feedFollow.UpdatedAt)
	fmt.Printf("User ID: %v\n", feedFollow.UserID)
	fmt.Printf("User Name: %v\n", feedFollow.UserName)
	fmt.Printf("Feed ID: %v\n", feedFollow.FeedID)
	fmt.Printf("Feed Name: %v\n", feedFollow.FeedName)

	return nil
}
