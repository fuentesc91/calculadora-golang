package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fuentesc91/calculadora/operadores"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("CALCULADORA, INGRESA LA OPERACION")
	scanner.Scan()
	operacion := scanner.Text()

	resultado := operadores.Operacion(operacion)

	fmt.Println(resultado)
}
