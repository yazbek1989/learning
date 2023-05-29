package main

import "fmt"

func logicos2() {
	//not (niega o cambia el resultado)
	fmt.Println(!false)

	//and (tienen que ser ambos true para devolver true)
	fmt.Println(false && true)

	//or (cualquiera que de true devuelve verdadero)
	fmt.Println(true || false)

	b := 2

	r := b == 2 || !(4 > b)
	fmt.Println(r)

}
