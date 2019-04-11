package main

import (
	"fmt"
	"github.com/gugafdsb1702/olago/fundamentos/funcoes/matematica"
)

func main(){
	resultado := matematica.Calculo(matematica.Multiplicador,2,2)
	fmt.Printf("O resultado do multiplicador foi: %d\r\n", resultado)
	resultado = matematica.Soma(3, 30)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Divisor, 12, 3)
	fmt.Printf("O resultado do divisor foi: %d\r\n", resultado)
	resultado, resto := matematica.DivisorComResto(12, 3)	
	fmt.Printf("O resultado do divisor foi: %d, e o resto é: %d\r\n", resultado, resto)
}	