package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/priyanshoon/feeder/internal/database"
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

func handlerCreateFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>\n", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	// get current user
	currentUser := s.config.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("cannot get the current user: %w\n", err)
	}

	// connect current user with feed and create new feed
	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	saveFeed, err := s.db.CreateFeed(context.Background(), feed)
	if err != nil {
		return fmt.Errorf("cannot create the feed: %w\n", err)
	}

	fmt.Println(saveFeed)

	return nil
}
