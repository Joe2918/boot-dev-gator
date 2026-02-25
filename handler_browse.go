package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/bootdotdev/boot-dev-gator/internal/database"
)

func browsePosts(s *state, cmd command, user database.User) error {
	var limitNum int64
	limitNum = 2
	if len(cmd.Args) == 1 {
		limitNum, _ = strconv.ParseInt(cmd.Args[0], 10, 32)
	}

	args := database.GetPostsParams{
		UserID: user.ID,
		Limit:  int32(limitNum),
	}

	posts, err := s.db.GetPosts(context.Background(), args)
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
