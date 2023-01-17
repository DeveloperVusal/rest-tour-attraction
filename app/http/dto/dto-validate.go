package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type DtoValidate struct{}

func (dtoeq *DtoValidate) Is(d any) bool {
	isValidate := true

	validate := validator.New()

	err := validate.Struct(d)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			isValidate = false
		}

		fmt.Println("------ List of tag fields with error ---------")

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.StructField())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println("---------------")
		}

		isValidate = false
	}

	return isValidate
}
