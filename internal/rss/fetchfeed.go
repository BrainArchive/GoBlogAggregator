package rss

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fixRss(rssFeed *RSSFeed) error {

	fixedTitle := html.UnescapeString(rssFeed.Channel.Title)
	fixedDesc := html.UnescapeString(rssFeed.Channel.Description)
	rssFeed.Channel.Title = fixedTitle
	rssFeed.Channel.Description = fixedDesc
	for _, rssItem := range rssFeed.Channel.Item {
		fixedTitle = html.UnescapeString(rssItem.Title)
		fixedDesc = html.UnescapeString(rssItem.Description)
		rssItem.Title = fixedTitle
		rssItem.Description = fixedDesc
	}

	return nil
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return nil, err
	}

	err = fixRss(&rssFeed)
	if err != nil {
		return nil, err
	}

	return &rssFeed, nil
}
