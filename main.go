package main

import (
	"fmt"
	"strings"
)

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
	inputArr := strings.Split(input, "")
	selected := ""
	selectCounter := 0
	control := false

	for i := 0; i < len(inputArr); i++ {
		if control {
			if selectCounter > 0 && selected == inputArr[i] {
				selectCounter--
				continue
			}

			if selectCounter == 0 {
				if inputArr[i] == ")" {
					return "Invalid"
				}

				control = false
				selected = inputArr[i]
				selectCounter++
				continue
			}

			if selectCounter > 0 {
				control = false
				selectCounter++
				selected = inputArr[i]
			}
		}

		if i == 0 {
			if inputArr[i] == ")" {
				return "Invalid"
			}

			selected = inputArr[i]
			selectCounter++
			continue
		}

		if i > 0 && selected == inputArr[i] {
			selectCounter++
			continue
		}

		if i > 0 && selected != inputArr[i] {
			selectCounter--
			selected = inputArr[i]
			control = true
			continue
		}
	}

	if selectCounter != 0 {
		return "Invalid"
	}

	return "Valid"

}
