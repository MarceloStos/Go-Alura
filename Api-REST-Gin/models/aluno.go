package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome      string `json:"nome" validate:"nonzero"`
	Idade     int    `json:"idade"`
	Curso     string `json:"curso"`
	Matricula string `json:"matricula" validate:"len=8, regexp=^[0-9]*$"`
}

func ValidaDadosDeAlunos(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}
	return nil
}
