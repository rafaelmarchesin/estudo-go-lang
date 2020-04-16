package main

import "fmt"

//declaração de variáveis
var x int = 10
var y int = 20

//declarando uma função
func add() int {
	return x + y
}

//executando a função principal
func main() {
	fmt.Println(add())
}
