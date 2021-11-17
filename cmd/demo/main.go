package main

import (
	"home/internal/demo/controller"
	"home/internal/demo/model"
	"home/pkg/utils/xgo"
	"home/pkg/xgorm"
	"log"
)

func main() {
	err := xgo.ParallelWithError(
		initDB,
		controller.Init,
	)()

	if err != nil {
		log.Fatalf("core run err: %v \n", err)
	}
}

// initDB connect to database
func initDB() error {
	db, err := xgorm.Build("mysql.test")
	if err != nil {
		return err
	}

	return model.Init(db)
}
