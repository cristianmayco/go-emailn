package main

import (
	"github.com/cristianmayco/go-emailn/internal/domain/campaign"
	"github.com/go-playground/validator/v10"
)

func main() {
	campaign := campaign.Campaign{}
	validate := validator.New()
	err := validate.Struct(campaign)

	if err == nil {
		println("Nenhum erro encontrado")
	} else {
		validationErros := err.(validator.ValidationErrors)
		for _, v := range validationErros {
			println(v.StructField() + " is invalid: " + v.Tag())
		}
	}
}
