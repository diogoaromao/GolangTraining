package main

import "fmt"

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Diogo"])
}

func changeMe(z map[string]int) {
	z["Diogo"] = 44
}
