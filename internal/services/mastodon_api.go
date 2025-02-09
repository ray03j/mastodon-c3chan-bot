package services

import (
	"log"
	"context"
	"os"
	"fmt"

	"github.com/mattn/go-mastodon"
)

var client *mastodon.Client

// Mastodonクライアントを作成する関数
func NewMastodonClient() *mastodon.Client {
	config := &mastodon.Config{
		Server:       os.Getenv("MASTODON_SERVER"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		AccessToken:  os.Getenv("ACCESS_TOKEN"),
	}

	return mastodon.NewClient(config)
}


func InitMastodonClient() {
	client = mastodon.NewClient(&mastodon.Config{
		Server:       os.Getenv("MASTODON_SERVER"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		AccessToken:  os.Getenv("ACCESS_TOKEN"),
	})
}

// ローカルタイムラインを監視
func LocalTimeLineListen(client *mastodon.Client) {
	fmt.Println("start streaming timeline")

	stream, err := client.StreamingPublic(context.Background(), false)
	if err != nil {
		log.Println("Error starting stream:", err)
		return
	}
	for event := range stream {
		switch status := event.(type) {
		case *mastodon.UpdateEvent:
			fmt.Printf("New post: %s, %s\n", status.Status.Content, status.Status.Account.Acct)
		}
	}
}

// トゥートを投稿
func PostToot(content string) {
	_, err := client.PostStatus(context.Background(), &mastodon.Toot{
		Status:     content,
		Visibility: "public",
	})
	if err != nil {
		log.Println("トゥートの投稿に失敗しました:", err)
	}
}

// トゥートにリプライ
func ReplyToot(content string) {
	_, err := client.PostStatus(context.Background(), &mastodon.Toot{
		Status:     content,
		Visibility: "public",
	})
	if err != nil {
		log.Println("リプライの投稿に失敗しました:", err)
	}
}
