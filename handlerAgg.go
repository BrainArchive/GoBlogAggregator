package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/brainarchive/goblogaggregator/internal/database"
	"github.com/brainarchive/goblogaggregator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {

	ticker := time.NewTicker(5 * time.Second)
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}

}

func scrapeFeeds(s *state) error {
	latestFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: time.Now(),
		ID:        latestFeed.ID,
	})

	url := latestFeed.Url
	rssFeed, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return err
	}
	fmt.Printf("%v", *rssFeed)
	fmt.Println()
	fmt.Println("====================")
	fmt.Printf("%v\n", latestFeed.Name)
	fmt.Printf("%v\n", latestFeed.Url)
	fmt.Printf("%v\n", latestFeed.UserID)
	user, err := s.db.GetUserFromId(context.Background(), latestFeed.UserID)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", user.Name)
	fmt.Println("====================")
	return nil
}
