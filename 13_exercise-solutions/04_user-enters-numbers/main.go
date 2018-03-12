package main

import "fmt"

func main() {
	var large int
	var small int

	fmt.Print("Large:")
	fmt.Scanln(&large)

	fmt.Print("Small:")
	fmt.Scanln(&small)

	fmt.Println(large / small)
}
