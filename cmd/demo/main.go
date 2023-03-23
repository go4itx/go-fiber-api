package main

import (
	"fmt"
	"home/internal/demo/controller"
	"home/internal/demo/model"
	"home/pkg/component/storage"
	"home/pkg/utils/xgo"
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	err := xgo.ParallelWithError(
		initDB,
		initScheduler,
		controller.Init,
	)()

	if err != nil {
		log.Fatalf("core run err: %v \n", err)
	}
}

// initDB connect to database
func initDB() error {
	gorm := storage.Gorm{}
	db, err := gorm.Build("mysql.test")
	if err != nil {
		return err
	}

	return model.Init(db)
}

// initScheduler start scheduler
func initScheduler() error {
	s := gocron.NewScheduler(time.Local)
	_, err := s.Every(10).Seconds().Do(func() {
		fmt.Println("hello world scheduler " + time.Now().Format("2006-01-02 15:04:05"))
	})

	s.StartAsync()
	return err
}
