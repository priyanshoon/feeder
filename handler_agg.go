package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("cannot fetch feed: %w\n", err)
	}

	fmt.Printf("Feed: %+v\n", feed)

	return nil
}
