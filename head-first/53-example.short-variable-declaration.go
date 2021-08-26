// 13-15: кратко объявляем переменные без слова var. Объявляем и тут же
// присваиваем им значения. Тип не указываем, потому что Go определяет
// тип, смотря на значение.
// 
// 17-20: выводим на экран переменные и строки (в последней функции ещё
// производим умноожение между двумя переменными).

package main

import "fmt"

func main() {
	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Damon Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}
