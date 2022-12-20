package main

import "fmt"

func relaciones2() {
	a := 4
	b := 5
	var r bool

	//igualdad
	r = a == b
	fmt.Printf("%d es igual que %d? %t \n", a, b, r)

	//distinto
	r = a != b
	fmt.Printf("%d es distinto que %d? %t \n", a, b, r)

	//mayor que
	r = a > b
	fmt.Printf("%d es mayor que que %d? %t \n", a, b, r)

	//menor que
	r = a < b
	fmt.Printf("%d es menor que que %d? %t \n", a, b, r)

	//mayor o igual que
	r = a >= b
	fmt.Printf("%d es mayor o igual que %d? %t \n", a, b, r)

	//menor o igual que
	r = a <= b
	fmt.Printf("%d es menor o igual que %d? %t \n", a, b, r)
}
