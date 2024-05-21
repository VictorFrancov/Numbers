package main

import "fmt"

func main() {
	fmt.Println("Primeira Parte")
	for numero := 1; numero <= 100; numero++ {
		if numero%3 == 0 {
			fmt.Println(numero)
		}
	}
	fmt.Println("----------------------------------")
	fmt.Println("//////////////////////////////////")
	fmt.Println("----------------------------------")
	fmt.Println("Segunda Parte")

	for numero := 1; numero <= 100; numero++ {
		if numero%3 == 0 {
			fmt.Println("PIN")
		} else if numero%5 == 0 {
			fmt.Println("PAN")

		}
	}
}
