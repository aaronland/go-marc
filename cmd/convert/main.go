package main

import (
	"context"
	"log"

	"github.com/aaronland/go-marc/v2/app/convert"
)

func main() {

	ctx := context.Background()
	err := convert.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run convert tool, %v", err)
	}
}
