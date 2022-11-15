package main

import "fmt"

// funcion para imprimirr nombre y texto segun recibido
func Funciones(nombre string) {
	fmt.Println("Hola, ", nombre)
}

//funcion para devolver valor entero e imprimirlo en el main
func Sumar(a, b int) int {
	total := a + b
	return total
}
