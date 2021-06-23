package main

import (
	"fmt"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

/*
* Escribir un código que permita validar si un string tiene correctamente balanceados los paréntesis.
* Ejemplo: "(()(()))" -> Valid
* "(())(()" -> Invalid
*
 */

// Validate imprime 'Valid' o 'Invalid' según el caso
//Validate("()") --> Valid
// Validate("(") --> Invalid
// Validate("(()(()))") --> Valid
// Validate("(())(()") --> Invalid
func main() {
	fmt.Println(Validate("()"))
	fmt.Println(Validate("("))
	fmt.Println(Validate("(()(()))"))
	fmt.Println(Validate("(())(()"))
}

func Validate(input string) string {
	//Write your code here

	var s Stack

	inputArr := strings.Split(input, "")

	for i := 0; i < len(inputArr); i++ {
		if inputArr[i] == "(" {
			s.Push(inputArr[i])
		} else {
			if s.IsEmpty() {
				return "Invalid"
			}

			b, status := s.Pop()

			if status {
				if b == "(" && inputArr[i] == ")" {
					continue
				}
			}

			return "Invalid"
		}
	}

	if s.IsEmpty() {
		return "Valid"
	}

	return "Invalid"
}
