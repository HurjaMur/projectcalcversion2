package main

import (
	"fmt"
	"time"

	"github.com/HurjaMur/projectcalcversion2/data"
	"github.com/HurjaMur/projectcalcversion2/orchestrator"
)

func main() {
	expression := "2+2*2" // Введите ваше выражение
	orch := orchestrator.NewOrchestrator(expression)
	exp, _ := orch.ParseAndProcessExpression()
	time.Sleep(1 * time.Second) // Введите сколько времени должно пройти до вычисления
	res := data.ExpressionResultMap[fmt.Sprint(expression)]
	fmt.Println("Ваше выражение:", "Ответ на ваше выражение: ", exp, res)
}
