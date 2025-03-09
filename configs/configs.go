package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type EnvVars struct {
	WebServerPort     int    `env:"WEB_SERVER_PORT"`
	MysqlRootPassword string `env:"MYSQL_ROOT_PASSWORD"`
	MysqlDatabase     string `env:"MYSQL_DATABASE"`
	MysqlUser         string `env:"MYSQL_USER"`
	MysqlPassword     string `env:"MYSQL_PASSWORD"`
	MysqlHost         string `env:"MYSQL_HOST"`
	MysqlPort         string `env:"MYSQL_PORT"`
}

func LoadEnv() *EnvVars {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env variables")
	}

	portStr := os.Getenv("WEB_SERVER_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid WEB_SERVER_PORT value: %s", portStr)
	}

	return &EnvVars{
		WebServerPort:     port,
		MysqlRootPassword: os.Getenv("MYSQL_ROOT_PASSWORD"),
		MysqlDatabase:     os.Getenv("MYSQL_DATABASE"),
		MysqlUser:         os.Getenv("MYSQL_USER"),
		MysqlPassword:     os.Getenv("MYSQL_PASSWORD"),
		MysqlHost:         os.Getenv("MYSQL_HOST"),
		MysqlPort:         os.Getenv("MYSQL_PORT"),
	}
}
