package agent

import (
	"regexp"
	"strconv"
	"strings"
)

func ProcessExpression(expression string) int {
	expression = strings.ReplaceAll(expression, " ", "")

	operands := regexp.MustCompile(`\d+`)
	operations := regexp.MustCompile(`[\+\-\*\/]`)
	numbers := operands.FindAllString(expression, -1)
	ops := operations.FindAllString(expression, -1)

	result, _ := calculateExpression(numbers, ops)

	return result
}

func calculateExpression(numbers []string, ops []string) (int, error) {
	result := 0
	currentOperation := "+"

	for i, num := range numbers {
		parsedNum, err := strconv.Atoi(num)
		if err != nil {
			return 0, err
		}

		if currentOperation == "+" {
			result += parsedNum
		} else if currentOperation == "-" {
			result -= parsedNum
		} else if currentOperation == "*" {
			result *= parsedNum
		} else if currentOperation == "/" {
			result /= parsedNum
		}

		// Установка операции для следующего числа, если есть
		if i < len(ops) {
			currentOperation = ops[i]
		}
	}

	return result, nil
}
