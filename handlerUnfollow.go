package main

import (
	"context"
	"errors"

	"github.com/brainarchive/goblogaggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return errors.New("usage unfollow <url>")
	}

	feed, err := s.db.FeedFromUrl(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	err = s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return err
	}

	return nil
}
