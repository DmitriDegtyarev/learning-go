// 12: преобразуем переменную length в Integer, чтобы можно было присвоить
// ей значение переменной другого типа. Когда идёт преобразование из float64
// в int, часть после точки округляется в меньшую сторону (отбрасывается).

package main

import "fmt"

func main() {
	var length float64 = 3.75
	var width int = 5
	width = int(length)
	fmt.Println(width)
}
