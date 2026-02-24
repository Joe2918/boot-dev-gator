package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/bootdotdev/boot-dev-gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time between request>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(timeBetweenRequests)
	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	return nil
}

func scrapeFeeds(s *state) error {
	feedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	fetchedTime := sql.NullTime{
		Time:  time.Now().UTC(),
		Valid: true,
	}
	args := database.MarkFeedFetchedParams{
		ID:            feedToFetch.ID,
		LastFetchedAt: fetchedTime,
	}

	err = s.db.MarkFeedFetched(context.Background(), args)
	if err != nil {
		return err
	}

	feed, err := fetchFeed(context.Background(), feedToFetch.Url)
	if err != nil {
		return err
	}

	fmt.Println(feedToFetch.Name)
	fmt.Println("=====================================")
	for _, item := range feed.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil
}
