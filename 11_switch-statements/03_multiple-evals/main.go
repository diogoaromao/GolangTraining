package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel", "Tim":
		fmt.Println("Wassup Daniel, or, err, Tim")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Have you no friends?")
	}
}
