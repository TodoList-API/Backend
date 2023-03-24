package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	MYSQL_DBUSER string
	MYSQL_DBPASS string
	MYSQL_DBHOST string
	MYSQL_DBPORT int
	MYSQL_DBNAME string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}
	isRead := true
	if val, found := os.LookupEnv("MYSQL_USER"); found {
		app.MYSQL_DBUSER = val
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_PASSWORD"); found {
		app.MYSQL_DBPASS = val
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_HOST"); found {
		app.MYSQL_DBHOST = val
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_PORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.MYSQL_DBPORT = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_DBNAME"); found {
		app.MYSQL_DBNAME = val
		isRead = false
	}

	if isRead {
		err := godotenv.Load("local.env")
		if err != nil {
			fmt.Println("Error saat baca env", err.Error())
			return nil
		}

		app.MYSQL_DBUSER = os.Getenv("MYSQL_USER")
		app.MYSQL_DBPASS = os.Getenv("MYSQL_PASSWORD")
		app.MYSQL_DBHOST = os.Getenv("MYSQL_HOST")
		readData := os.Getenv("MYSQL_PORT")
		app.MYSQL_DBPORT, err = strconv.Atoi(readData)
		if err != nil {
			fmt.Println("Error saat convert MYSQL_DBPORT", err.Error())
			return nil
		}
		app.MYSQL_DBNAME = os.Getenv("MYSQL_DBNAME")
	}

	return &app
}
