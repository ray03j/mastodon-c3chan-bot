package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mattn/go-mastodon"
)

func main() {
	config := &mastodon.Config{
		Server:       os.Getenv("MASTODON_SERVER"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		AccessToken:  os.Getenv("ACCESS_TOKEN"),
	}

	// Create the client
	c := mastodon.NewClient(config)

	// Post a toot
	// finalText := "test"
	// visibility := "public"

	toot := mastodon.Toot{
		Status:     "test",
		Visibility: "public",
	}
	fmt.Println("Posting toot:", toot)

	post, err := c.PostStatus(context.Background(), &toot)
	if err != nil {
		log.Fatalf("%#v\n", err)
	}

	fmt.Println("My new post is:", post)
}