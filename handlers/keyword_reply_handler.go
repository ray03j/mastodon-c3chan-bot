package handlers

import (
	"fmt"
	"strings"

	"mastodon-c3chan-bot/internal/services"
)

// キーワードが含まれていたらリプライを送信
func HandleKeywordReply(content string, username string) {
	keywords := []string{"こんにちは", "おはよう", "ありがとう"}

	for _, keyword := range keywords {
		if strings.Contains(content, keyword) {
			reply := fmt.Sprintf("@%s こんにちは！", username)
			services.ReplyToot(reply)
			return
		}
	}
}
