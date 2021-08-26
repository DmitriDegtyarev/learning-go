// 10: импортируем пакет reflect, чтобы можно было использовать функцию TypeOf,
// которая проверяет тип значения переменной.
//
// 14-17: выводим на экран результаты исполнения функции TypeOf

package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Hello, Go!"))
}
