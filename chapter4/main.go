package main

import "fmt"

var con = "Hello World"

func main() {
	// 1 вид присвоения
	var x string = "Hello World"
	var y string
	x = "Hello World"
	y = `Hello World`
	fmt.Println(x)
	fmt.Println(y)
	// 2 вид присвоения
	var a, b string
	a = "Hello World"
	b = `Hello World`
	fmt.Println(a)
	fmt.Println(b)
	// 3 вид присвоения
	var c = "Hello World"
	fmt.Println(c)
	// конкатенация
	var z string
	z = "first "
	fmt.Println(z)
	z = z + "second"
	fmt.Println(z)
	z = "first "
	z += "second"
	fmt.Println(z)
	// сравнение строк
	var q = "hello"
	var w = "world"
	fmt.Println(q == w)
	var e = "hello"
	var r = "hello"
	fmt.Println(e == r)
	// присвоение переменной
	t := "hello"
	fmt.Println(t)
	// глобальная переменная
	fmt.Println(con)
	// константа
	const cont = "Hello World"
	fmt.Println(cont)
	// вызов другой функции
	f()
	input()
}

func f() {
	// несколько переменных
	var (
		a = 5
		b = 10
	)
	// несколько констант
	const (
		c = 11
		d = 12
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(con)
}

func input() {
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
