// 11: преобразуем переменную width в float64, чтобы можно было присвоить
// ей значение переменной другого типа.

package main

import "fmt"

func main() {
	var length float64 = 1.2
	var width int = 2
	length = float64(width)
	fmt.Println(length)
}
