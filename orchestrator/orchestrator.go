package orchestrator

import (
	"github.com/HurjaMur/projectcalcversion2/agent"
	"github.com/HurjaMur/projectcalcversion2/data"
)

type Base struct {
	Expression string
}

func NewOrchestrator(expression string) Base {
	return Base{
		Expression: expression,
	}
}

func isExpressionValid(expr string) bool {
	if len(expr) == 0 {
		return false
	}

	return isValidExpression(expr, 0)
}

func isValidExpression(expr string, index int) bool {
	if index == len(expr) {
		return true
	}

	if expr[index] >= '0' && expr[index] <= '9' {
		return isValidExpression(expr, index+1)
	} else if expr[index] == '+' || expr[index] == '-' || expr[index] == '*' || expr[index] == '/' {
		return isValidExpression(expr, index+1)
	}

	return false
}

func (b Base) ParseAndProcessExpression() (string, error) {
	expression := b.Expression
	if !isExpressionValid(expression) { //Проверяем на валидность
		return "400. Выражение невалидно", nil
	}

	if !isValidExpression(expression, 0) { //Проверяем на валидность
		return "400. Выражение невалидно", nil
	}
	if isValidExpression(expression, 0) && isExpressionValid(expression) {
		result := agent.ProcessExpression(expression) //Считаем с помощью агента
		data.SaveToDatabase(expression, result)       //Сохраняем в базу данных
	}

	return "200. Выражение успешно принято, распаршено и принято к обработке", nil
}
