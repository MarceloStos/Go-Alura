package database

import (
	"log"

	"github.com/MarceloStos/Go-Alura/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Err error
)

func ConectacomBancoDeDados() {
	stringDeConexao := "host=127.0.0.1 user=myuser password=123456 dbname=mydatabase port=5434 sslmode=disable"
	DB, Err = gorm.Open(postgres.Open(stringDeConexao))
	if Err != nil {
		log.Panic("Não foi possível conectar com o banco de dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
