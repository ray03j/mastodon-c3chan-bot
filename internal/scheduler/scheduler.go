package scheduler

import (
	"github.com/go-co-op/gocron"
	"time"
)

// スケジューラを開始する関数
func StartScheduler() {
	scheduler := gocron.NewScheduler(time.Local)

	// 1時間ごとのタスクを登録
	scheduler.Every(2).Hour().Do(TaskBoost)

	// 毎日のタスクを登録
	scheduler.Every(1).Day().Do(ExecuteOnChime)

	// 毎月のタスクを登録
	scheduler.Every(1).Month().Do(MonthChange)

	// スケジューラを開始
	scheduler.StartAsync()
}
