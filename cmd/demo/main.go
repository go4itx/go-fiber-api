package main

import (
	"home/internal/demo/controller"
	"home/internal/demo/service"
	"home/pkg/utils/xgo"
	"log"
)

func main() {
	err := xgo.ParallelWithError(
		controller.Init,
		service.Init,
	)()

	if err != nil {
		log.Fatalf("core run err: %v \n", err)
	}
}
