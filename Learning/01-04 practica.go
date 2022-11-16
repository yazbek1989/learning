package main

import "fmt"

func testsumdiv() {
	var a int
	var b int
	fmt.Println("ingrese el primer valor: ")
	fmt.Scan(&a)
	fmt.Println("ingrese el segundo valor: ")
	fmt.Scan(&b)
	fmt.Println("el resultado de su suma es: ", sumas(a, b))
	coc, res := dividir(a, b)
	fmt.Println("el cociente es: ", coc)
	fmt.Println("el residuo es: ", res)
}
func sumas(a, b int) int {
	return a + b
}
func dividir(a, b int) (coc int, res int) {
	coc = a / b
	res = a % b
	return coc, res
}

func testtaxes() {
	var a float64
	fmt.Println("ingrese el valor de compra del producto: ")
	fmt.Scan(&a)
	tax, total := impuesto(a)
	fmt.Printf("el IGV es %v y sumado al producto es: %v ", tax, total)
}

func impuesto(a float64) (tax float64, total float64) {
	b := 0.18
	tax = b * a
	total = tax + a
	return tax, total
}
