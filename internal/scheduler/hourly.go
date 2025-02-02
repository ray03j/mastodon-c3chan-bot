package scheduler

import (
	"time"

	"mastodon-c3chan-bot/internal/services"
)

// #今日やること #明日やること
func TaskBoost() {
	ticker := time.NewTicker(2 * time.Hour)
	go func() {
		for {
			<-ticker.C
			services.PostToot("これは2時間ごとの定期トゥートです！")
		}
	}()
}
