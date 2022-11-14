package main

import "fmt"

func Format() {
	hola := "Hola"
	mundo := "Mundo"
	fmt.Println(hola, ", ", mundo)

	nombre := "Iván"
	edad := 33
	//%T es para ver que tipo de dato es
	fmt.Printf("nombre: %T \n", nombre)
	//%s es para string y %d es para entero
	fmt.Printf("Hola, %s y su edad es %d \n", nombre, edad)
	//en printf %v es para cualquier tipo de variable
	fmt.Printf("Hola, %v y su edad es %v \n", nombre, edad)
	//sprintf es para guardar un mensaje en una variable
	mensaje := fmt.Sprintf("Hola, %v y su edad es %v", nombre, edad)
	fmt.Println(mensaje)

}
