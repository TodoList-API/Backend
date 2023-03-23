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
	if val, found := os.LookupEnv("MYSQL_DBUSER"); found {
		app.MYSQL_DBUSER = val
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_DBPASS"); found {
		app.MYSQL_DBPASS = val
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_DBHOST"); found {
		app.MYSQL_DBHOST = val
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_DBPORT"); found {
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

		app.MYSQL_DBUSER = os.Getenv("MYSQL_DBUSER")
		app.MYSQL_DBPASS = os.Getenv("MYSQL_DBPASS")
		app.MYSQL_DBHOST = os.Getenv("MYSQL_DBHOST")
		readData := os.Getenv("MYSQL_DBPORT")
		app.MYSQL_DBPORT, err = strconv.Atoi(readData)
		if err != nil {
			fmt.Println("Error saat convert MYSQL_DBPORT", err.Error())
			return nil
		}
		app.MYSQL_DBNAME = os.Getenv("MYSQL_DBNAME")
	}

	return &app
}
