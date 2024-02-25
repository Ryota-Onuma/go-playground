package context

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type UserIDKeyType struct{}
type RoleKeyType struct{}

var UserIDKey = UserIDKeyType{}
var RoleKey = RoleKeyType{}

func RunCounter(ctx context.Context, index int) {
	checkDeadLine(ctx)
	checkValue(ctx)
	go echoHelloWorld(ctx, index+3)
	printLogForCounter(index, "カウントを出力するGoroutineを起動するよ")
	var count int
LOOP:
	for {
		select {
		case <-ctx.Done():
			errCheck(ctx.Err())
			printLogForCounter(index, "キャンセルのシグナルが来たから終了するよ")
			break LOOP
		default:
			count++
			printLogForCounter(index, strconv.Itoa(count))
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// 孫GoroutineにContentが伝播することを確認するための関数
func echoHelloWorld(ctx context.Context, index int) {
	checkDeadLine(ctx)
	checkValue(ctx)
	printLogForCounter(index, "Hello, World!")
	select {
	case <-ctx.Done():
		errCheck(ctx.Err())
		printLogForCounter(index, "キャンセルのシグナルが来たから終了するよ。ちなみに孫Goroutineだよ")
	}
}

func printLogForCounter(index int, text string) {
	fmt.Printf("%v番目のGoroutine: %v\n", index, text)
}

func errCheck(err error) {
	if errors.Is(err, context.Canceled) {
		fmt.Println("キャンセルされたよ", err.Error())
	} else if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("デッドラインを超えたよ", err.Error())
	} else {
		fmt.Println(err)
	}
}

func checkDeadLine(ctx context.Context) {
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println("デッドラインは", deadline.String(), "だよ")
	} else {
		fmt.Println("デッドラインはないよ")
	}
}

func checkValue(ctx context.Context) {
	if userID, ok := ctx.Value(UserIDKey).(string); ok {
		fmt.Println("userIDは", userID, "だよ")
	} else {
		fmt.Println("userIDはないよ")
	}
	if role, ok := ctx.Value(RoleKey).(string); ok {
		fmt.Println("roleは", role, "だよ")
	} else {
		fmt.Println("roleはないよ")
	}
}
