package main

import (
	"context"
	"github.com/Coldwws/chat_practice/internal/di"
	"log"
)

func main() {
	ctx := context.Background()
	a, err := di.NewApp(ctx)
	if err != nil {
		log.Fatalf("Failed to init app:%v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("Failed to run:%v", err)
	}
}
