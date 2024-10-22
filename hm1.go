package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Calc(expression string) (float64, error) {
	var numbers []float64
	var signs []rune

	calculate := func(sig rune) error {
		if len(numbers) < 2 {
			return fmt.Errorf("Недостаточно слагаемых")
		}
		a := numbers[len(numbers)-2]
		b := numbers[len(numbers)-1]
		numbers = numbers[:len(numbers)-2]

		var res float64
		switch sig {
		case '+':
			res = a + b
		case '-':
			res = a - b
		case '*':
			res = a * b
		case '/':
			if b == 0 {
				return fmt.Errorf("Нельзя делить на ноль")
			}
			res = a / b
		default:
			return fmt.Errorf("Неизвестный оператор: %c", sig)
		}
		numbers = append(numbers, res)
		return nil
	}

	priority := func(sig rune) int {
		switch sig {
		case '+', '-':
			return 1
		case '*', '/':
			return 2
		default:
			return 0
		}
	}

	var number strings.Builder

	for _, char := range expression {
		if unicode.IsDigit(char) || char == '.' {
			number.WriteRune(char)
		} else {
			if number.Len() > 0 {
				num, err := strconv.ParseFloat(number.String(), 64)
				if err != nil {
					return 0, err
				}
				numbers = append(numbers, num)
				number.Reset()
			}

			if char == ' ' {
				continue
			}

			if char == '(' {
				signs = append(signs, char)
			} else if char == ')' {
				for len(signs) > 0 && signs[len(signs)-1] != '(' {
					if err := calculate(signs[len(signs)-1]); err != nil {
						return 0, nil
					}
					signs = signs[:len(signs)-1]
				}
				if len(signs) == 0 {
					return 0, fmt.Errorf("Дана закрывающая скобка, но нет открывающей")
				}
				signs = signs[:len(signs)-1]
			} else {
				for len(signs) > 0 && priority(signs[len(signs)-1]) >= priority(char) {
					if err := calculate(signs[len(signs)-1]); err != nil {
						return 0, err
					}
					signs = signs[:len(signs)-1]
				}
				signs = append(signs, char)
			}
		}
	}
	if number.Len() > 0 {
		num, err := strconv.ParseFloat(number.String(), 64)
		if err != nil {
			return 0, err
		}
		numbers = append(numbers, num)
	}
	for len(signs) > 0 {
		if err := calculate(signs[len(signs)-1]); err != nil {
			return 0, err
		}
		signs = signs[:len(signs)-1]
	}
	if len(numbers) != 1 {
		return 0, fmt.Errorf("Неправильное выражение")
	}
	return numbers[0], nil
}

func main() {

}
