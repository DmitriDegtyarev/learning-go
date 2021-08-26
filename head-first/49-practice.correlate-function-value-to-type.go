// 13-19: объявляем переменные, указываем их тип и значение.
//
// 21-27: выводим на экран типы переменных.

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int = 25
	var b bool = true
	var c float64 = 5.2
	var d int = 1
	var e bool = false
	var f float64 = 1.0
	var g string = "hello"

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))
	fmt.Println(reflect.TypeOf(e))
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(g))
}
