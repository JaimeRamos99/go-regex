package main
import (
  "regexp"
)
func eval_pattern(pattern string) string{

  match, _ := regexp.MatchString("^[a-z]{3}[0-9]{3}$", pattern)

  if match {
    return "SI puede ser una placa"
  }
  
  return "NO puede ser una placa"

}
