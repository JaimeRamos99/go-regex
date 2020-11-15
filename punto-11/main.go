package main
import (
"fmt"
"os"
)

func main() {
  //Getting command line arguments
  args := os.Args
  args_length := len(args)

  //Validating that there is at list 1 car plate to evaluate
  if args_length <= 1 {
    fmt.Println("¡Se necesita por lo menos una posible placa para verificar!")
    return
  }

  //Getting only the args we are interested in
  car_plates := os.Args[1:]

  //creates an instance of color blue
  colorYellow := "\033[33m"
  colorReset := "\033[0m"

  //validates if each possible plate, meets the required pattern
  for i, s := range car_plates {
    fmt.Println(i, " ", string(colorYellow), s, " ", string(colorReset), eval_pattern(s))
  }

}
