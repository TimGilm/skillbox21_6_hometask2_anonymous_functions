// Задание 2. Анонимные функции
// Что нужно сделать
// Напишите функцию, которая на вход принимает функцию вида A func (int, int) int,
// а внутри оборачивает и вызывает её при выходе (через defer).
// Вызовите эту функцию с тремя разными анонимными функциями A.
// Тела функций могут быть любыми, но главное, чтобы все три
// выполняли разное действие.
package main

import "fmt"

func f(x, y int, a func(x, y int) int) int {
	fmt.Println(a(x, y))
	return a(x, y)
}
func main() {
	defer f(2, 3, func(x, y int) int { return x + y })
	defer f(2, 3, func(x, y int) int { return x - y })
	defer f(2, 3, func(x, y int) int { return x * y })
}
