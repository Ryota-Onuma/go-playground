package main

import (
	"context"
	"log"
	"log/slog"
	"os"
)

// SEE: https://zenn.dev/88888888_kota/articles/7e97ff874083cf#slog%E3%81%AE%E7%9B%AE%E6%A8%99
func main() {
	trySlog()
}

func trySlog() {
	// デフォルトのロガーはプレーンテキストになる
	// デフォルトではログレベルがInfo以上のログのみが出力される
	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")

	logLevel := new(slog.LevelVar)
	ops := slog.HandlerOptions{
		Level: logLevel,
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &ops))

	// デフォルトのロガーをJSONHandlerに変更
	slog.SetDefault(logger)

	// JSON形式のログになる
	slog.Info("Info message")
	log.Println("Info message")

	// slog.Loggerではなく、log.Loggerを使ってログを出力する
	handler := slog.NewJSONHandler(os.Stdout, nil)
	infoLog := slog.NewLogLogger(handler, slog.LevelInfo)
	WarningLog := slog.NewLogLogger(handler, slog.LevelWarn)
	errorLog := slog.NewLogLogger(handler, slog.LevelError)

	infoLog.Println("Something has occurred...")
	WarningLog.Println("Warning!...")
	errorLog.Println("error has occurred")

	// slog.Attrを使ってログを出力する
	slog.LogAttrs(context.Background(),
		slog.LevelInfo,
		"structured logger message",
		slog.String("method", "GET"),
		slog.Int("status", 200),
		slog.Group(
			"group1", slog.Group(
				"nested1", slog.String(
					"key1", "value1",
				),
			),
		),
	)

	// ログレベルがDebug上が表示されるようにする
	logLevel.Set(slog.LevelDebug)

	slog.Debug("Debug message")
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")

	// ログレベルがError以上が表示されるようにする
	logLevel.Set(slog.LevelError)
	slog.Info("Info message")
	slog.Warn("Warning message")
	slog.Error("Error message")
}
