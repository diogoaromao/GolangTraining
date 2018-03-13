package main

import "fmt"

func main() {
	name := "Diogo"
	changeMe(name)
	fmt.Println(name)
}

func changeMe(z string) {
	fmt.Println(z)
	z = "Andre"
}
