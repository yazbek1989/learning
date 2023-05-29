package main

import "fmt"

// funcion para imprimirr nombre y texto segun recibido
func funciones(nombre string) {
	fmt.Println("Hola, ", nombre)
}

//funcion para devolver valor entero e imprimirlo en el main
func sumar(a, b int) int {
	total := a + b
	return total
}
