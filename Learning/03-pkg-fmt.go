package main

import "fmt"

func Format() {
	hola := "Hola"
	mundo := "Mundo"
	fmt.Println(hola, ", ", mundo)

	nombre := "Iván"
	edad := 33
	fmt.Printf("Hola, %s y su edad es %d \n", nombre, edad)
	fmt.Printf("Hola, %v y su edad es %v \n", nombre, edad)
}
