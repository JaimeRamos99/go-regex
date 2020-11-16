package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	// Se toma por comando el string a ingresar y el entero X correspondiente a la cantidad de decimales
	args := os.Args
	argsLength := len(args)

	if argsLength <= 2 {
		fmt.Println("¡Se necesita por lo menos un numero con n decimales para verificar!")

	} else {
		input := os.Args[1]
		n := os.Args[2]

		if i, err := strconv.Atoi(n); err == nil {

			// Expresión regular

			cadena := fmt.Sprintf(`[+-]?([0-9]*[.])?[0-9]+`)
			verificarCadena := regexp.MustCompile(cadena)

			if verificarCadena.FindString(input) == "" {
				fmt.Printf("No es un número real")
				return
			}

			if i < 0 {
				fmt.Printf("La cantidad de decimales ingresada no puede ser negativa")
				return
			}
			decimal := fmt.Sprintf(`^-?[0-9]+(?:\.\d{%d,%d})?$`, i, i)

			verificarDecimal := regexp.MustCompile(decimal)

			if verificarDecimal.FindString(input) == "" {
				fmt.Printf("No tiene la cantidad exacta de %d decimales", i)
			} else {

				if len(input) == 1 && i != 0 {
					fmt.Printf("Es un número real pero no tiene la cantidad exacta de %d decimales", i)
				} else {
					if i == 1 {
						fmt.Printf(" %v Es un numero real con %d decimal", verificarDecimal.FindString(input), i)
					} else {
						fmt.Printf(" %v Es un numero real con %d decimales", verificarDecimal.FindString(input), i)
					}

				}

			}

		}
	}

}
