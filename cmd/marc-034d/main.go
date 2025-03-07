// marc-034d is a web application for converting MARC 034 strings in to bounding boxes (formatted as GeoJSON).
package main

import (
	"context"
	"log"
	
	"github.com/aaronland/go-marc/v2/app/server"
)

func main() {

	ctx := context.Background()
	err := server.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run server, %v", err)
	}
}
