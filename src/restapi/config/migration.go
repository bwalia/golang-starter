package config

import (
	"golang-starter/src/restapi/models"
)

func Migrate() {
	InitDb().AutoMigrate(&models.Book{})
}
