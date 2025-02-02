package utils

import (
	"log"
	"fmt"
	"os"
)

// ログレベルの定義
const (
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
)

// ログの出力先
var (
	logFile *os.File
	logger  *log.Logger
)

// 初期化関数
func InitLogger(logFilePath string) error {
	var err error

	// ログファイルを開く（存在しなければ作成する）
	logFile, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// ロガーを作成（標準出力とファイルの両方に出力）
	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}

// INFO レベルのログを出力
func Info(message string, args ...interface{}) {
	logger.Printf("[%s] %s", INFO, formatMessage(message, args...))
}

// WARN レベルのログを出力
func Warn(message string, args ...interface{}) {
	logger.Printf("[%s] %s", WARN, formatMessage(message, args...))
}

// ERROR レベルのログを出力
func Error(message string, args ...interface{}) {
	logger.Printf("[%s] %s", ERROR, formatMessage(message, args...))
}

// ログファイルを閉じる
func CloseLogger() {
	if logFile != nil {
		logFile.Close()
	}
}

// メッセージのフォーマット
func formatMessage(message string, args ...interface{}) string {
	if len(args) > 0 {
		return fmt.Sprintf(message, args...)
	}
	return message
}
