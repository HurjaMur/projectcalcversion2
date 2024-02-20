package main

import (
	"fmt"
	"time"

	"github.com/HurjaMur/projectcalcversion2/data"
	"github.com/HurjaMur/projectcalcversion2/orchestrator"
)

func main() {
	var expression string
	fmt.Print("Введите ваше выражение")
	x, _ := fmt.Scan(&expression)                    // Введите ваше выражение
	orch := orchestrator.NewOrchestrator(expression) //Добавляем наше выражение в оркестратор
	exp, _ := orch.ParseAndProcessExpression()
	time.Sleep(1 * time.Second) // Введите сколько времени должно пройти до вычисления
	x++                         //Просто юзаем, чтобы не было ошибок в компиляторе
	res := data.ExpressionResultMap[fmt.Sprint(expression)]
	fmt.Printf("Корректность вашего выражения: %v, Ответ на ваше выражение: %v\n", exp, res)
}
