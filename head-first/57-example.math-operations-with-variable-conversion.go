// 12-13: преобразуем переменную width в float64, чтобы можно было произвести
// математические операции между переменными (они могут производиться только
// между переменными одного типа). Выводим результат на экран.

package main

import "fmt"

func main() {
	var length float64 = 1.2
	var width int = 2
	fmt.Println("Area is", length*float64(width))
	fmt.Println("length > width?", length > float64(width))
}
