package orchestrator

import (
	"math/rand"
	"regexp"

	"github.com/HurjaMur/projectcalcversion2/agent"
	"github.com/HurjaMur/projectcalcversion2/data"
)

type Base struct {
	Expression string
	ID         int
}

func GenerateID() int {
	return rand.Intn(1000000) // Генерим случайный айди
}

func NewOrchestrator(expression string) Base {
	return Base{
		Expression: expression,
		ID:         GenerateID(),
	}
}

func (b Base) ParseAndProcessExpression() (string, error) {
	expression := b.Expression

	// Валидность выражения
	validChars := regexp.MustCompile(`^[0-9\+\-\*\/\s()]+$`)

	// Проверка валидности выражения
	if !validChars.MatchString(expression) {
		return "400. Выражение невалидно", nil
	}

	result := agent.ProcessExpression(expression)

	// Сохранение результата в базу данных
	data.SaveToDatabase(b.ID, expression, result)

	return "200. Выражение успешно принято, распаршено и принято к обработке", nil
}
