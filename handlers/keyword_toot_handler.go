package handlers

import (
	"fmt"
	"strings"

	"mastodon-c3chan-bot/internal/services"
)

// キーワードが含まれていたら新しくトゥート
func HandleKeywordToot(content string) {
	keywords := []string{"天気", "ニュース", "おすすめ"}

	for _, keyword := range keywords {
		if strings.Contains(content, keyword) {
			toot := fmt.Sprintf("キーワード '%s' に関連する情報をお届けします！", keyword)
			services.PostToot(toot)
			return
		}
	}
}
