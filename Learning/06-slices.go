package main

import "fmt"

func slices() {
	numeros := []int{1, 2, 3}

	//agregar datos a slice
	numeros = append(numeros, 4)
	numeros = append(numeros, 5)

	fmt.Println(numeros, len(numeros))

	//sacar subslice del slice
	subNumeros := numeros[:2]
	numeros[0] = 100

	fmt.Println(numeros)
	fmt.Println(subNumeros)

	//punteros,
	//longitud,
	//capacidad
	meses := []string{
		"enero",
		"febrero",
		"marzo",
	}
	fmt.Printf("len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")

	//el slice se volvio a generar con el append, entonces el slice anterior se borro y se creo uno nuevo en otra referencia de memoria y de esa forma van creciendo
	//es como un arrey que se va sustituyendo para crecer
	fmt.Printf("len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)
}
