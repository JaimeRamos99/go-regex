package main
import (
  "regexp"
)
func eval_pattern(pattern string) string{

  match, _ := regexp.MatchString("([a-z][a-z][a-z][0-9][0-9][0-9])", pattern)

  if match {
    return "SI puede ser una placa"
  }
  
  return "NO puede ser una placa"

}
