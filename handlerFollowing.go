package main

import (
	"context"
	"fmt"

	"github.com/brainarchive/goblogaggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	for _, feedFollow := range feedFollows {
		feed, err := s.db.FeedFromID(context.Background(), feedFollow.FeedID)
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", feed.Name)
	}

	return nil
}
