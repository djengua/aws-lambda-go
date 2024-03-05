package tools

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/djengua/aws-lambda-go/models"
)

func GetDate() string {
	t := time.Now()

	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func LoadConfig() (config models.ConfigInfo, err error) {

	_, exist := os.LookupEnv("DB_DRIVER")
	if !exist {
		err = errors.New("no se configuro env")
		return
	}
	config.DBDriver = os.Getenv("DB_DRIVER")

	_, exist = os.LookupEnv("DB_SOURCE")
	if !exist {
		err = errors.New("no se configuro env")
		return
	}
	config.DBSource = os.Getenv("DB_SOURCE")

	return
}
