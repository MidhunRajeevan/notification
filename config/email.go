package config

import (
	"log"
	"os"
)

type EmailConfig struct {
	UserName string
	Password string
	Sender   string
}

const (
	emailUserName = "USER_NAME"
	emailPassword = "PASSWORD"
	Sender        = "SENDER"
)

var Econfig EmailConfig

func EmalInit() {
	if value, ok := os.LookupEnv(emailUserName); ok {
		Econfig.UserName = value
	} else {
		log.Println("The environment variable USER_NAME is not set")
		os.Exit(1)
	}
}
