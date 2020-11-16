package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	// Se valida si se ingresa algún argumento
	args := os.Args
	argsLength := len(args)

	if argsLength <= 2 {
		fmt.Println("¡Se necesita por lo menos un numero con n decimales para verificar!")

	} else {
		// Se toma por comando el string a ingresar y el entero X correspondiente a la cantidad de decimales
		input := os.Args[1]
		n := os.Args[2]

		if i, err := strconv.Atoi(n); err == nil {

			// Expresión regular que verifica si lo ingresado corresponde a una cadena o un número. Si es una cadena (no númerica), debe salirse del programa

			cadena := fmt.Sprintf(`[+-]?([0-9]*[.])?[0-9]+`)
			verificarCadena := regexp.MustCompile(cadena)

			if verificarCadena.FindString(input) == "" {
				fmt.Printf("No es un número real")
				return
			}

			// Verificar que la cantidad de decimales no sea negativa
			if i < 0 {
				fmt.Printf("La cantidad de decimales ingresada no puede ser negativa")
				return
			}

			// Expresión regular principal, donde se verifica que el número ingresado tiene i cantidad de decimales.

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
