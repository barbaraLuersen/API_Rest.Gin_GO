package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero, regexp=^[a-zA-Z\s]*$"` //nonzero valida se é zero, nulo ou vazio
	CPF  string `json:"cpf" validate:"len=14, regexp=^[0-9.-]*$"`      //len valida o número máximo de caracteres
	RG   string `json:"rg" validate:"len=12, regexp=^[0-9.-]*$"`       //regexp valida o conteúdo da string, aceitando apenas aquilo que foi setado entre colchetes []
}

// Função para validar o conteúdo passado. Se der tudo certo não retorna nada
func ValidaDados(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		return err
	}
	return nil
}
