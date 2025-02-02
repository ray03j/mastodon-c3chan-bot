package scheduler

import (
	"time"
	"mastodon-c3chan-bot/internal/services"
)

// 授業のお知らせ
func ExecuteOnChime() {
	// 毎日決まった時間に
	// 定刻をどこかのファイルで設定する
	ticker := time.NewTicker(2 * time.Hour)
	go func() {
		for {
			<-ticker.C
			services.PostToot("これは2時間ごとの定期トゥートです！")
		}
	}()
}