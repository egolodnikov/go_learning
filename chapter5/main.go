package main

import "fmt"

func main() {
	// цикл for and if else
	i := 1
	for i <= 10 {
		fmt.Println(i)
		if i%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
		i += 1
	}
	// цикла for and if else if else
	j := 1
	for j <= 10 {
		if j%2 == 0 {
			fmt.Println("divided by 2")
		} else if j%3 == 0 {
			fmt.Println("divide by 3")
		} else {
			fmt.Println("not divided")
		}
		j += 1
	}
	// for and switch
	x := 1
	for x <= 10 {
		switch x % 5 {
		case 0:
			fmt.Println("divide on 5")
		case 1:
			fmt.Println("divided by 5")
		default:
			fmt.Println("not divide")

		}
		x += 1
	}

}
