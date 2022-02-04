package main

import "fmt"

func main() {
	// целые и вещественные числа
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("1.0 + 1.0 = ", 1.0+1.0)
	// строки
	fmt.Println("Hello World")
	fmt.Println(`Hello World`)
	// экранирует символы
	fmt.Println("Hello\tWorld")
	// не экранирует
	fmt.Println(`Hello\tWorld`)
	// операции со строкам
	fmt.Println(len("Hello World"))
	fmt.Println(len(`Hello World`))
	fmt.Println("Hello World"[1])
	fmt.Println(`Hello World`[1])
	fmt.Println("Hello " + "World")
	fmt.Println(`Hello ` + `World`)
	// булевые типы
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
