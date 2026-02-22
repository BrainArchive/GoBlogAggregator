package main

import (
	"context"
	"fmt"
	"strconv"
)

func handlerBrowse(s *state, cmd command) error {
	var limit int64
	if len(cmd.args) >= 1 {
		var err error
		limit, err = strconv.ParseInt(cmd.args[0], 10, 32)
		if err != nil {
			return err
		}
	} else {
		limit = 2
	}

	fmt.Printf("BROWSING %v POSTS \n", limit)

	posts, err := s.db.GetPosts(context.Background(), int32(limit))
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Printf("TITLE:   %v\n", post.Title)
		fmt.Printf("UPDATED: %v\n", post.UpdatedAt)
		fmt.Printf("CREATED: %v\n", post.CreatedAt)
		fmt.Printf("URL:     %v\n", post.Url)
		fmt.Printf("DESC:    %v\n", post.Description)
		fmt.Printf("FEEDID:  %v\n", post.FeedID)
	}

	return nil
}

// type Post struct {
// 	ID          uuid.UUID
// 	CreatedAt   time.Time
// 	UpdatedAt   time.Time
// 	Title       string
// 	Url         string
// 	Description sql.NullString
// 	PublishedAt sql.NullTime
// 	FeedID      uuid.UUID
// }
