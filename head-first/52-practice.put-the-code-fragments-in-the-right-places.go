// 13-14: объявляем переменные (указываем имена и тип)
//
// 16-18: выводим на экран строки и значения переменных (последняя операция
// ещё и одну переменную вычитает из другой)

package main

import (
	"fmt"
)

func main() {
	var originalCount int = 10
	var eatenCount int = 4

	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left.")
}
