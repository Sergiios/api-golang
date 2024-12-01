package utils_test

import (
	"api-golang/internal/utils"
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestFormatValidationErrors_WithValidationErrors(t *testing.T) {
	validate := validator.New()

	// Struct de teste com tags de validação
	type TestStruct struct {
		Name string `validate:"required"`
		Age  int    `validate:"min=18"`
	}

	// Instância inválida
	testData := TestStruct{
		Name: "",
		Age:  15,
	}

	// Validação da struct
	err := validate.Struct(testData)

	// Confirma que há erros de validação
	assert.Error(t, err)

	// Formata os erros usando a função
	formattedErr := utils.FormatValidationErrors(err)

	// Confirma que o erro foi formatado corretamente
	expectedErrorMessage := "Name failed validation on required. Age failed validation on min. "
	assert.EqualError(t, formattedErr, expectedErrorMessage)
}

func TestFormatValidationErrors_WithNonValidationError(t *testing.T) {
	// Erro genérico
	nonValidationError := errors.New("generic error")

	// Formata o erro usando a função
	formattedErr := utils.FormatValidationErrors(nonValidationError)

	// Confirma que o erro original é retornado sem modificação
	assert.EqualError(t, formattedErr, "generic error")
}

func TestFormatValidationErrors_WithNilError(t *testing.T) {
	// Passa nil como erro
	formattedErr := utils.FormatValidationErrors(nil)

	// Confirma que nil é retornado
	assert.Nil(t, formattedErr)
}
