// Задание 2. Анонимные функции
// Что нужно сделать
// Напишите функцию, которая на вход принимает функцию вида x func (int, int) int,
// а внутри оборачивает и вызывает её при выходе (через defer).
// Вызовите эту функцию с тремя разными анонимными функциями x.
// Тела функций могут быть любыми, но главное, чтобы все три
// выполняли разное действие.
package main

import "fmt"

func main() {
	fmt.Println(calculate(func(x, y int) int { return x + y }, 2, 3))
	fmt.Println(calculate(func(x, y int) int { return x * y }, 3, 2))
	fmt.Println(calculate(func(x, y int) int { return x / y }, 3, 3))
}

func calculate(a func(int, int) int, x int, y int) (result int) {
	defer func() {
		result = a(x, y)
	}()
	result = 0 // эта строка демонстрирует, что функция в defer меняет значение при выходе из функции,
	// эту строку можно удалить
	return
}
