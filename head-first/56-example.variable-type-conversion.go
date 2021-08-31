// 14: преобразуем переменную myInt из Integer в Float64 и выводим на экран тип
// этой переменной.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int = 2
	fmt.Println(reflect.TypeOf(myInt))
	fmt.Println(reflect.TypeOf(float64(myInt)))
}
