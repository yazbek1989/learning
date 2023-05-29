package main

import "fmt"

func arrays() {
	var numeros [5]int
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	fmt.Println(numeros)
	fmt.Println(numeros[1])

	//arrays con valores
	nombres := [3]string{"Alex", "Roel", "Juan"}
	fmt.Println(nombres)

	//array con longitud definida en el momento, posteriormente no se puede ampliar el arrey
	colores := [...]string{
		"rojo",
		"verde",
		"negro",
		"azul",
	}
	fmt.Println(colores, len(colores))

	monedas := [...]string{
		0: "dolar",
		2: "Soles",
		3: "pesos",
		5: "euro",
	}
	fmt.Println(monedas, len(monedas))

	//sub arrey (saca uno antes del indice final) [0:3] o [:3] o [1:]
	subMoneda := monedas[:4]
	fmt.Println(subMoneda, len(subMoneda))
}
