package main

import (
	"fmt"
	"os"
)

func main() {
	var first, second float64
	var do string
	fmt.Println("Введите первое число: ")
	fmt.Fscan(os.Stdin, &first)
	fmt.Println("Введите действие: ")
	fmt.Fscan(os.Stdin, &do)
	fmt.Println("Введите второе число: ")
	fmt.Fscan(os.Stdin, &second)
	if do == "+" {
		fmt.Println(first + second)
	} else if do == "-" {
		fmt.Println(first - second)
	} else if do == "*" {
		fmt.Println(first * second)
	} else if do == "/" {
		if second == 0 {
			fmt.Println("На ноль делить нельзя")
			return
		}
		fmt.Println(first / second)
	} else {
		fmt.Println("Неизвестное действие")
	}
}
