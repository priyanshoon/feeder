package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	items := RSSFeed{}
	err = xml.Unmarshal(data, &items)
	if err != nil {
		return nil, err
	}

	items.Channel.Title = html.UnescapeString(items.Channel.Title)
	items.Channel.Description = html.UnescapeString(items.Channel.Description)
	for i := range items.Channel.Item {
		items.Channel.Item[i].Title = html.UnescapeString(items.Channel.Item[i].Title)
		items.Channel.Item[i].Description = html.UnescapeString(items.Channel.Item[i].Description)
	}

	return &items, nil
}
