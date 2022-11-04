package database

import (
	"log"

	"github.com/felipepnascimento/challenge-bravo-flp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	databaseUrl := "host=postgres13 user=postgres password=postgres dbname=challenge_bravo_dev port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(databaseUrl))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Currency{})
	DB.AutoMigrate(&models.Conversion{})
}