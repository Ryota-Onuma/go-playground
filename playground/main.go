package main

import (
	"context"
	"fmt"
	pcontext "go-playground/context"
	"time"
)

func main() {
	playContext()
}

func playContext() {
	playCancel()
	playWithTimeout()
	playWithDeadline()
	playValue()
}

func playCancel() {
	fmt.Println("キャンセルで遊んでみる")
	ctx, cancel := context.WithCancel(context.Background())
	for i := 1; i <= 3; i++ {
		go pcontext.RunCounter(ctx, i)
	}
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
	fmt.Println("キャンセルしたよ")
	fmt.Println("")
}

func playWithTimeout() {
	fmt.Println("タイムアウト(WithTimeout)で遊んでみる")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	for i := 1; i <= 3; i++ {
		go pcontext.RunCounter(ctx, i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("タイムアウト(WithTimeout)したよ")
	fmt.Println("")
}

func playWithDeadline() {
	fmt.Println("デッドライン(WithDeadline)で遊んでみる")
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	for i := 1; i <= 3; i++ {
		go pcontext.RunCounter(ctx, i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("デッドライン(WithDeadline)したよ")
	fmt.Println("")
}

func playValue() {
	fmt.Println("Valueで遊んでみる")
	ctx := context.WithValue(context.Background(), pcontext.UserIDKey, "hogehogehoge_userID")
	ctx = context.WithValue(ctx, pcontext.RoleKey, "admin")
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	for i := 1; i <= 3; i++ {
		go pcontext.RunCounter(ctx, i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Valueで遊んで、タイムアウトしたよ")
}
