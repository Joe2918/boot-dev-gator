package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Joe2918/boot-dev-gator/internal/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
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

func scrapeFeeds(s *state) {
	feedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
		return
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
		log.Printf("Couldn't mark feed %s fetched: %v", feedToFetch.Name, err)
		return
	}

	feed, err := fetchFeed(context.Background(), feedToFetch.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feedToFetch.Name, err)
		return
	}

	fmt.Println(feedToFetch.Name)
	fmt.Println("=====================================")
	for _, item := range feed.Channel.Item {
		description := sql.NullString{
			String: item.Description,
			Valid:  true,
		}

		pubTime, err := time.Parse("time.RFC3339", item.PubDate)

		publishedTime := sql.NullTime{
			Time:  pubTime,
			Valid: true,
		}

		args := database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Url:         item.Link,
			Description: description,
			PublishedAt: publishedTime,
			FeedID:      feedToFetch.ID,
		}
		err = s.db.CreatePost(context.Background(), args)
		if err != nil {
			if pqErr, ok := err.(*pq.Error); ok {
				if pqErr.Code == "23505" {
					continue
				}
			}
			log.Printf("Couldn't create post: %v", err)
		}
	}

	log.Printf("Feed %s collected, %v posts found", feedToFetch.Name, len(feed.Channel.Item))
}
