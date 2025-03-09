package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"tasks-api/configs"
)

func New(envs *configs.EnvVars) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", envs.MysqlUser, envs.MysqlPassword, envs.MysqlHost, envs.MysqlPort, envs.MysqlDatabase))
	if err != nil {
		log.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	return db
}
