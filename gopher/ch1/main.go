package main

import (
	"context"

	"github.com/rema424/hexample"
)

func main() {
	ctx := context.Background()
	msg := "こんばんは"
	hexample.Run(ctx, msg)
}
