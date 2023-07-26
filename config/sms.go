package config

import (
	"log"
	"os"
)

type SMSConfig struct {
	SmppServer string
	SmppPort   string
	UserName   string
	Password   string
}

const (
	SmppServer  = "SMPP_SERVER"
	SmppPort    = "SMPP_PORT"
	smsUserName = "SMS_USER_NAME"
	smsPassword = "PASSWORD"
)

func SmsInit() {
	if value, ok := os.LookupEnv(smsUserName); ok {
		Econfig.UserName = value
	} else {
		log.Println("The environment variable SMS_USER_NAME is not set")
		os.Exit(1)
	}

}
