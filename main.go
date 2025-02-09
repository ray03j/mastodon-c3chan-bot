package main

import (
	// "mastodon-c3chan-bot/internal/services"
	// "mastodon-c3chan-bot/internal/utils"
	// "mastodon-c3chan-bot/tests"
	// "log"

	"time"
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Start Mastodon C3-chan Bot")
	fmt.Println(time.Local)
	// c := services.NewMastodonClient()
	// tests.PostTestToot(c)

    // スケジューラを Goroutine で実行
    go startScheduler()

	// メイン関数が終了しないようにする
	var wg sync.WaitGroup
	wg.Add(1) // 永続的にブロックする
	wg.Wait()

	// スケジューラ開始
	// scheduler.StartScheduler()

	// メイン関数が終了しないようにする
	select {}

	// ログファイルの初期化
	// err := utils.InitLogger("bot.log")
	// if err != nil {
	// 	log.Fatalf("ログの初期化に失敗しました: %v", err)
	// }
	// defer utils.CloseLogger()

	// utils.Info("Botが起動しました")
	// utils.Warn("APIのレスポンスが遅いです")
	// utils.Error("APIのリクエストに失敗しました: %s", "エラー内容")
}

// スケジューラのダミー関数
func startScheduler() {
    for {
        fmt.Println("Tooting at:", time.Now().Format("2006-01-02 15:04:05"))
        time.Sleep(2 * time.Hour) // 2時間ごとに実行
    }
}
