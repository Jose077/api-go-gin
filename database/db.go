package database

import (
	"log"

	"github.com/jose077/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringConexao := "host=172.19.0.2 user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(stringConexao))
	log.Println(DB, "eeeeeeeee")
	if err != nil {
		log.Panic("Erro ao conectar com o db!", err.Error())
	}

	log.Print("db success!")

	DB.AutoMigrate(&models.Aluno{})

}
