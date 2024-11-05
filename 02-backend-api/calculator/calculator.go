package calculator

import (
	"errors"

	"github.com/basola21/02-backend-api/utils"
)

type Numbers struct {
	Number1 string `json:"number1"`
	Number2 string `json:"number2"`
}

type Results struct {
	Result float64 `json:"result"`
}

func CalculateTwoNumbers(operation string, numbers Numbers) (Results, error) {
	number1, err := utils.CastString(numbers.Number1)
	number2, err := utils.CastString(numbers.Number2)
	if err != nil {
		return Results{}, err
	}

	var result Results

	switch operation {
	case "add":
		result = Results{Result: number1 + number2}
	case "subtract":
		result = Results{Result: number1 - number2}
	case "multiply":
		result = Results{Result: number1 * number2}
	case "divide":
		if number2 == 0 {
			return Results{}, errors.New("can not divide by 0")
		}
		result = Results{Result: number1 / number2}
	default:
		return Results{}, errors.New("please ender a valid operation")
	}
	return result, nil
}
