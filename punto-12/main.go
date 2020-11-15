package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	// Se toma por comando el string a ingresar y el entero X correspondiente a la cantidad de decimales
	input := os.Args[2]
	n := os.Args[3]

	if i, err := strconv.Atoi(n); err == nil {

		// Expresión regular
		response := fmt.Sprintf(`^[0-9]+(?:\.\d{%d,%d})?$`, i, i)

		re := regexp.MustCompile(response)

		if re.FindString(input) == "" {
			fmt.Printf("No es un número real o no tiene la cantidad exacta de %d decimales", i)
		} else {
			if i == 1 {
				fmt.Printf(" %v Es un numero real con %d decimal", re.FindString(input), i)
			} else {
				fmt.Printf(" %v Es un numero real con %d decimales", re.FindString(input), i) // True
			}

		}

	}

}