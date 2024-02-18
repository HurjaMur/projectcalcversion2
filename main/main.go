package main

import (
	"fmt"
	"time"

	"github.com/HurjaMur/projectcalcversion2/data"
	"github.com/HurjaMur/projectcalcversion2/orchestrator"
)

func main() {
	expression := "2+2*2"                                // Введите ваше выражение
	orch := orchestrator.NewOrchestrator(expression)
	exp, _ := orch.ParseAndProcessExpression()
	time.Sleep(10 * time.Second)                         // Введите сколько времени должно пройти до вычисления
	res := data.ExpressionResultMap[fmt.Sprint(orch.ID)] // используйте fmt.Sprint для преобразования int в строку
	fmt.Printf("\nВаше выражение: %s, Ответ на ваше выражение: %d", exp, res.Result)
}
