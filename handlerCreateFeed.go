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

func handlerCreateFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return errors.New("usage gator feed <name> <url>")
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	})
	if err != nil {
		fmt.Printf("Error when creating feed: %v\n", err)
		os.Exit(1)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
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

	fmt.Printf("feed generated %v\n", feed)

	return nil
}
