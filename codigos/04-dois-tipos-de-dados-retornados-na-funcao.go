package main

import "fmt"

func swap(x string, y int) (string, int) {
	return x,y
}

func main() {
	a, b := swap("hello", 10)
	fmt.Println(a, b)
	
	c := "Posso usar o := para declarar uma variável"
	fmt.Println(c)
}

