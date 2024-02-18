package main

import (
	"fmt"
	"time"

	"github.com/HurjaMur/projectcalcversion2/data"
	"github.com/HurjaMur/projectcalcversion2/orchestrator"
)

func main() {
	expression := "2+2*2" // Введите ваше выражение
	orch := orchestrator.NewOrchestrator(expression) //Добавляем наше выражение в оркестратор
	exp, _ := orch.ParseAndProcessExpression()
	time.Sleep(1 * time.Second) // Введите сколько времени должно пройти до вычисления
	res := data.ExpressionResultMap[fmt.Sprint(expression)]
	fmt.Print("Корректность вашего выражения,", "ответ на ваше выражение соотвественно ", exp, res)
}
