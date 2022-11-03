package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		i float64 = 3.141592654
		j string  = "test"
	)
	j = strconv.FormatFloat(i, 'E', -1, 64)
	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", j, j)
}
