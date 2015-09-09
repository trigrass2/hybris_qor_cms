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
	Host              string
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)

	Env = envOrPanic("HYBRIS_ENV", false)

	MySQLHost = envOrPanic("HYBRIS_MYSQL_PORT_3306_TCP_ADDR", false)
	MySQLPort = envOrPanic("HYBRIS_MYSQL_PORT_3306_TCP_PORT", false)
	MySQLRootPassword = envOrPanic("HYBRIS_MYSQL_ENV_MYSQL_ROOT_PASSWORD", true)

	MySQLDatabase = envOrPanic("HYBRIS_MYSQL_DATABASE", false)
	Verbose = (envOrPanic("HYBRIS_VERBOSE", true) != "")
	Host = envOrPanic("HYBRIS_HOST", false)
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
