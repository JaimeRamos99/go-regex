package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)

	fmt.Print("Enter number: ")
	var n int

	fmt.Scan(&n)
	response := fmt.Sprintf(`^[0-9]+(?:\.\d{%d,%d})?$`, n, n)
	//cadena = fmt.Sprintf(`[-]?\d[\d]*[\.](\d{%d})`)
	re := regexp.MustCompile(response)

	fmt.Printf("Contiene  %v\n", re.FindString(input)) // True

}
