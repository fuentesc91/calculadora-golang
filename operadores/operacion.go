package operadores

import (
	"fmt"
	"strconv"
	"strings"
)

// Operacion realiza la operacion sin importar que se usen varios signos
func Operacion(operacion string) int {
	operadores, nuevaOperacion := tomarOperadores(operacion)
	valores := strings.Split(nuevaOperacion, " ")
	resultado, error := strconv.Atoi(valores[0])

	if error != nil {
		fmt.Println("Error", error, "Debes de usar solo números enteros y operadores válidos")
	}

	for i := 0; i < len(operadores); i++ {
		num, error := strconv.Atoi(valores[i+1])
		if error != nil {
			fmt.Println("Error", error, "Debes de usar solo números enteros y operadores válidos")
		}
		if operadores[i] == "+" {
			resultado += num
		} else if operadores[i] == "-" {
			resultado -= num
		} else if operadores[i] == "*" {
			resultado *= num
		} else if operadores[i] == "/" {
			resultado /= num
		} else {
			fmt.Println("Error  Debes de usar solo números enteros y operadores válidos")
		}
	}
	return resultado
}

func tomarOperadores(operacion string) ([]string, string) {
	var (
		index      int
		operadores []string
	)
	for true {
		index = strings.IndexAny(operacion, "+-*/")
		if index == -1 {
			break
		}
		operadores = append(operadores, operacion[index:index+1])
		operacion = strings.Replace(operacion, operacion[index:index+1], " ", 1)
	}
	if strings.ContainsAny(operacion, "'%'qwertyuiopasdfghjklñzxcvbnm_.;,:") {
		fmt.Println("Error Tu operación contiene un caracter invalido")
	}
	return operadores, operacion
}
