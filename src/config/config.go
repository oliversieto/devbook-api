package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DataSourceName = ""
	ApiPort        = ""
	SecretKey      []byte
)

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ApiPort = os.Getenv("API_PORT")

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

	DataSourceName = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PHRASE"),
		os.Getenv("DB_NAME"))

}
