package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MidhunRajeevan/notification/app"
	"github.com/MidhunRajeevan/notification/config"
)

func main() {

	config.InitApp()
	config.EmalInit()
	config.SmsInit()

	http.HandleFunc("/notification", app.Notification)

	log.Println("server starting on ", config.Appcon.ListerningPort)
	url := fmt.Sprintf(":%s", config.Appcon.ListerningPort)
	http.ListenAndServe(url, nil)

}
