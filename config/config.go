package config

import (
	"os"
	"log"
	"strconv"
)



var Env = struct {
	DatabaseURL      string
	DatabaseName     string
	DatabaseUser     string
	ListenPort     	 string
	DatabasePassword string
}{
	DatabaseURL:      "DATABASE_URL",
	DatabaseName:     "DATABASE_NAME",
	DatabaseUser:     "DATABASE_USER",
	ListenPort:       "LISTEN_PORT",
	DatabasePassword: "DATABASE_PASSWORD",
}

func Get(envVariable string) string {
	return os.Getenv(envVariable)
}

func GetFailIfEmpty(envVariable string) string {
	val := os.Getenv(envVariable)
	if val == "" {
		log.Fatal("required environment variable is not defined: ", envVariable)
	}
	return val
}


func GetDatabaseUser() string {
	return GetFailIfEmpty(Env.DatabaseUser)
}

func GetListenPort() string {
	raw := GetFailIfEmpty(Env.ListenPort)
	port, err := strconv.ParseUint(raw, 10, 32)
	if err != nil || port > 65534 {
		log.Fatal("Wrong listen port value")
	}

	return raw
}


func GetDatabaseName() string {
	return GetFailIfEmpty(Env.DatabaseName)
}


func GetDatabasePassword() string {
	return GetFailIfEmpty(Env.DatabasePassword)
}


func GetDatabaseUrl() string {
	return GetFailIfEmpty(Env.DatabaseURL)
}
