package main

import (
	"os"
	"fmt"
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

//function that asks the user how many decimals is going to use the regexp
func ask_precision_decimals() int {

		fmt.Println("Digite el numero de decimales de precisión que deben tener los números:")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
	  line := scanner.Text()
		n_args :=len(strings.Fields(line))

		if n_args == 1{

			match_int, err := regexp.MatchString("^[0-9]+$", line)
			if err != nil || !match_int {
				fmt.Println("¡No digitó un entero correcto!")
				return -1
			}

			n_int, _ := strconv.Atoi(line)
			return n_int

		}
		fmt.Println("¡No digitó un entero correcto!")
		return -1

}

//function that asks the users for the numbers to validate if meets the regexp rules
func ask_numbers_to_test() []string {

	fmt.Println("Digite los números a probar, separados por un espacio.")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
  line := scanner.Text()

	cases := strings.Fields(line)

	return cases
}

//function that iterates over every number given by the user and determine if meet the pattern
func eval_pattern(numbers []string, decimals int) {

	colorYellow := "\033[33m"
	colorReset := "\033[0m"

	if decimals == 0 {

		for _, s := range numbers {

			match,_ := regexp.MatchString("^[0-9]+$", s)
			if match {
				fmt.Println(string(colorYellow), s, string(colorReset),"SI es un decimal de", decimals, "decimales de precisión")
				continue
			}

			fmt.Println(string(colorYellow), s, string(colorReset), "NO es un decimal de", decimals, "decimales de precisión")

		}

	} else {

		reg_exp := fmt.Sprintf(`^\-?[0-9]+\.[0-9]{%d}$`, decimals)

		for _, s := range numbers {

			match, _ := regexp.MatchString(reg_exp, s)

			if match {
				fmt.Println(string(colorYellow), s, string(colorReset), "SI es un decimal de", decimals, "decimales de precisión")
				continue
			}

			fmt.Println(string(colorYellow), s, string(colorReset), "No es un decimal de", decimals, "decimales de precisión")

		}

	}

}

func main() {

	decimal_precision := ask_precision_decimals()
	//only positive numbers accepted
	if decimal_precision >= 0 {
		array_numbers_to_test := ask_numbers_to_test()
		eval_pattern(array_numbers_to_test, decimal_precision)
	}


}
