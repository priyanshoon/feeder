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

	// follow feed
	feedfollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	createFeedFollow, err := s.db.CreateFeedFollow(context.Background(), feedfollow)
	if err != nil {
		return fmt.Errorf("cannot create feed follow row: %w\n", err)
	}

	fmt.Println(createFeedFollow)
	fmt.Println(saveFeed)

	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("cannot fetch the feeds: %w\n", err)
	}

	for i := range feeds {
		fmt.Printf("%s %s %s\n", feeds[i].Name, feeds[i].Url, feeds[i].Name_2)
	}
	return nil
}

func handlerFeedFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name> <url>\n", cmd.Name)
	}

	url := cmd.Args[0]

	// get feed by url
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("cannot fetch the url %w\n", err)
	}

	// get current user
	currentUser := s.config.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("error fetching current user : %w\n", err)
	}

	feedfollow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	createFeedFollow, err := s.db.CreateFeedFollow(context.Background(), feedfollow)
	if err != nil {
		return fmt.Errorf("cannot create feed follow row: %w\n", err)
	}

	fmt.Println(createFeedFollow.FeedName, createFeedFollow.UserName)

	return nil
}

func handlerFeedFollowing(s *state, cmd command) error {
	// get current user
	currentUser := s.config.CurrentUserName
	user, err := s.db.GetUser(context.Background(), currentUser)
	if err != nil {
		return fmt.Errorf("cannot fetch user in database: %w\n", err)
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("cannot get the following feed of user %s, %w\n", currentUser, err)
	}

	for i := range following {
		fmt.Println(following[i].FeedName)
	}

	return nil
}
