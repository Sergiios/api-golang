package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

// Função para formatar erros de validação
func FormatValidationErrors(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		// Concatena os erros em uma única mensagem
		errorMessages := ""
		for _, e := range validationErrors {
			errorMessages += e.Field() + " failed validation on " + e.Tag() + ". "
		}
		return errors.New(errorMessages)
	}
	return err
}
