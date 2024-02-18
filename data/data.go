package data

import (
	"sync"
)

type ID struct {
	Id     int
	Result int
}

var Mutex = &sync.Mutex{}
var ExpressionResultMap = make(map[string]ID)

func SaveToDatabase(id int, expression string, result int) {
	Mutex.Lock()
	defer Mutex.Unlock()
	ExpressionResultMap[expression] = ID{Id: id, Result: result}
}
