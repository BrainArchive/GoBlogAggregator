package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/brainarchive/goblogaggregator/internal/database"
	"github.com/brainarchive/goblogaggregator/internal/rss"
	"github.com/google/uuid"
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
	for _, rssItem := range rssFeed.Channel.Item {
		timePublished, err := time.Parse(time.RFC1123Z, rssItem.PubDate)
		if err != nil {
			fmt.Printf("error parsing time %v\n", err)
		}
		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     rssItem.Title,
			Url:       rssItem.Link,
			Description: sql.NullString{
				String: rssItem.Description,
				Valid:  true,
			},
			PublishedAt: sql.NullTime{
				Time:  timePublished,
				Valid: true,
			},
			FeedID: latestFeed.ID,
		})
		if err != nil {
			fmt.Printf("error creating post: %v\n", err)
		}
	}
	fmt.Println()
	fmt.Println("====================")
	fmt.Printf("%v\n", latestFeed.Name)
	fmt.Printf("%v\n", latestFeed.Url)
	fmt.Printf("%v\n", latestFeed.UserID)
	_, err = s.db.GetUserFromId(context.Background(), latestFeed.UserID)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", user.Name)
	fmt.Println("====================")

	return nil
}
