package main

import "fmt"

func print_monda(x int32) (int32, error) {
	if x < 0 {
		return 0, fmt.Errorf("Numero negativo")
	}
	return 12, nil
}

func main() {
	fmt.Println("Hello lets go")

	dato, err := print_monda(-1)

	if err != nil {
		panic("Se Jodio")
	}

	fmt.Println(dato)
}
