package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Curso string `json:"curso"`
}
