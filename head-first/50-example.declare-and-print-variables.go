// 16-18: объявляем переменные (указываем имя и тип. Если у нескольких
// переменных одинаковый тип, то можно перечислить их через запятую.
//
// 20-22: присваиваем значения переменным (для каждой переменной должно быть
// присвоено значение именно того типа, который мы указали в объявлении).
// То есть строковой переменной нельзя присвоить число и наоборот.
//
// 24-27: выводим на экран строки, переменные (в последней строке ещё и
// умножение между двумя переменными выполняем).

package main

import "fmt"

func main() {
	var quantity int
	var length, width float64
	var customerName string

	quantity = 4
	length, width = 1.2, 2.4
	customerName = "Damon Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters")
}
