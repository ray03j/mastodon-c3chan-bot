package main

import (
	"mastodon-c3chan-bot/internal/services"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start Mastodon C3-chan Bot")
	fmt.Println(time.Local)
	client := services.NewMastodonClient()

	go services.LocalTimeLineListen(client)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
}
