package service

import "log"

var (
	IM    = newIMService()
	Admin = newAdminService()
)

func init() {
	go IM.receiver()
	go IM.sender()
}

// errorHandler
func errorHandler(title string) {
	if err := recover(); err != nil {
		log.Println(title, ":", err)
	}
}
