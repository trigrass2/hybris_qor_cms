package config

import (
	"log"
	"os"
	"strings"
)

// this is configured from env variables
var (
	Env               string
	MySQLHost         string
	MySQLPort         string
	MySQLDatabase     string
	MySQLRootPassword string
	Verbose           bool
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

	Env = envOrPanic("DEVICEM_ENV", false)

	MySQLHost = envOrPanic("DEVICEM_MYSQL_PORT_3306_TCP_ADDR", false)
	MySQLPort = envOrPanic("DEVICEM_MYSQL_PORT_3306_TCP_PORT", false)
	MySQLRootPassword = envOrPanic("DEVICEM_MYSQL_ENV_MYSQL_ROOT_PASSWORD", true)

	MySQLDatabase = envOrPanic("DEVICEM_MYSQL_DATABASE", false)
	Verbose = (envOrPanic("DEVICEM_VERBOSE", true) != "")
}

func envOrPanic(key string, allowEmpty bool) (r string) {
	r = os.Getenv(key)
	if r == "" && !allowEmpty {
		panic("env " + key + " is not set")
	}
	logValue := r
	if strings.Contains(logValue, "PASSWORD") || strings.Contains(logValue, "SECRET") {
		logValue = "<HIDDEN>"
	}
	log.Printf("Configure: %s = %s\n", key, logValue)
	return
}
