// 33 строка - указывается основной пакет main, всё, что ниже - относится к
// этому пакету. Вообще, вроде как во всех программах, написанных на Go,
// в самом начале указывается именно пакет main.
// Пока с этими пакетами не совсем ясно. Для чего main вначале постоянно
// указывать - не понятно.
// Если бы я вначале вместо main указал пакет fmt, то получается его
// и импортировать бы не пришлось, лишнюю строку кода написал.
// Не знаю, может дальше об этом будет поподробнее объяснено, но пока всё туго.
//
// Что за пакеты, для чего нужны, за что каждый из них отвечает и тд, хз.
// Вроде как пакеты - это что-то наподобие библиотек. Они просто
// предоставляют функции и остальные возможности, которые можно будет
// использовать в программе. Пакеты, наверное находятся все в папке, куда
// установлен Go. Чтобы для каждой проги не подключать все доступные пакеты,
// которые есть в папке, надо подключать только те, которые будут
// использоваться в программе (это значительно ускоряет программу. Для этого
// и нужен import).
//
// 35 строка - мпортируется пакет fmt, чтобы можно было использовать функцию из
// этого пакета далее в программе.
//
// 37 строка - указывается функция и тело функции. Видимо основная
// функция тут из пакета main, а fmt - это так сказать подфункция (функция,
// которая в свою очередь относится к другой, "главной" функции).
//
// 38 строка - указывается вначале название пакета, к которому относится
// функция, потом указывается название функции, в скобках указывается значение
// функции, двойные ковычки указывают на то, что это строка.
//
// Понять бы для чего package и main нужны. Почему в самом начале именно этот
// пакет указывается, да и почему бы его просто не указать в import-е, хз.

package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")
}
