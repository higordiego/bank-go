package main

import "fmt"

type contaCorrent struct {
	titular      string
	numerAgencia int
	numeroConta  int
	saldo        float64
}

func main() {

	conta1 := contaCorrent{"Higor Diego", 489, 123456, 125.20}
	fmt.Println(conta1)
}
