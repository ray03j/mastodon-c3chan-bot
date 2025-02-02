package scheduler

import (
	"time"
	"mastodon-c3chan-bot/internal/services"
)

// 毎月1日の0時にトゥート
func MonthChange() {
	go func() {
		for {
			now := time.Now()
			next := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location())

			// 次の1日までスリープ
			time.Sleep(time.Until(next))
			services.PostToot("新しい月が始まりました！")
		}
	}()
}
