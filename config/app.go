package config

import "os"

type Appconfig struct {
	ListerningPort string
}

var Appcon Appconfig

const (
	appListerningPort = "APP_LISTERNING_PORT"
	defaultPort       = "9091"
)

func InitApp() {
	if value, ok := os.LookupEnv(appListerningPort); ok {
		Appcon.ListerningPort = value
	} else {
		Appcon.ListerningPort = defaultPort
	}
}
