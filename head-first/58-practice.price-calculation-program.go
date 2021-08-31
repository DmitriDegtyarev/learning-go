// 27: присваиваем переменной tax новое значение, которое будет равняться
// умножению переменных price и taxRate между собой. Чтобы эти переменные
// можно было перемножить, приводим их к одному типу (либо к float64 либо к
// integer. В данном случае price переводим в float64, а не taxRate в int, так
// как в этом случае вместо 0.08 будет 0).
//
// С остальными преобразованиями ниже - всё то же самое. Преобразуем именно в
// float64, а не в int, чтобы в случае, если мы поменяем значения переменных,
// у нас программа не похерилась (все значения не округлились, вдруг я везде
// захочу использовать числа с плавающей точкой, а не только в переменной
// taxRate).
//
// 35: преобразуем availableFunds в float64 и сравниваем наш бюджет с итоговой
// ценой. Если наш бюджет равен или больше цены, тогда на экране появится
// строка "Within budget?, true". Если же наш бюджет меньше цены, тогда вместо
// true будет false.

package main

import "fmt"

func main() {
	var price int = 100
	fmt.Println("Price is", price, "dollars.")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")

	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", total, "dollars.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "dollars available.")
	fmt.Println("Within budget?", float64(availableFunds) >= total)
}
