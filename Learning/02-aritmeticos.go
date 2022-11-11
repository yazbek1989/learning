package main

import "fmt"

func Aritmeticos() {
	a := 21
	b := 10

	fmt.Println("el valor de a es:", a)
	fmt.Println("el valor de b es:", b)

	//suma
	result := a + b
	fmt.Println("suma:", result)

	//resta
	result = a - b
	fmt.Println("resta:", result)

	//multiplicacion
	result = a * b
	fmt.Println("multiplicacion:", result)

	//division
	result = a / b
	fmt.Println("division:", result)

	//division decimal
	var div float64 = 3.0 / 2.0
	fmt.Println("division:", div)

	//modulo
	result = a % b
	fmt.Println("modulo:", result)

}
