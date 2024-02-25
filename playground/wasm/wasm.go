package main

import (
	"bytes"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/yuin/goldmark"
)

func setUp() {
	c := make(chan struct{}, 0)
	js.Global().Set("markdownToHtml", js.FuncOf(markdownToHtml))
	<-c
}

func markdownToHtml(_ js.Value, args []js.Value) any {
	in := args[0].String()
	out := toHtml(in)
	return out
}

func toHtml(in string) string {
	inBuf := []byte(in)
	var outBuf bytes.Buffer
	if err := goldmark.Convert(inBuf, &outBuf); err != nil {
		fmt.Printf("fail to convert markdown: %v\n", err)
		return in
	}

	out := strings.ReplaceAll(outBuf.String(), "\n", "")
	return out
}
