package main

import (
	"context"
	"fmt"
)

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("cannot scrape %w\n", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("cannot mark the feed %w\n", err)
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("cannot fetch the feed %w\n", err)
	}

	for _, v := range feed.Channel.Item {
		fmt.Println(v.Title)
	}

	return nil
}
