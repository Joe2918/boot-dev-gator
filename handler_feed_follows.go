package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bootdotdev/boot-dev-gator/internal/database"
	"github.com/google/uuid"
)

func handlerFeedFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return err
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	args := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), args)
	if err != nil {
		return err
	}
	fmt.Printf("Feed: %s\nCurrent user: %s\n", feedFollow.FeedName, feedFollow.UserName)
	return nil
}
