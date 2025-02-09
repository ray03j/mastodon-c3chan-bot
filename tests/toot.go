package tests

import (
	"context"
	"fmt"
	"log"
	"github.com/mattn/go-mastodon"
)

// テスト用に投稿を行う関数
func PostTestToot(c *mastodon.Client) {
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
