package main

import (
	"home/internal/im/controller"
	"home/pkg/xgo"
	"log"
)

func main() {
	err := xgo.ParallelWithError(
		controller.Init,
	)()

	if err != nil {
		log.Fatalf("core run err: %v \n", err)
	}
}
