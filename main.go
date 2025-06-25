package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const version = "1.0.0"

// CalculateRPN は逆ポーランド記法の文字列を計算します。
func CalculateRPN(input string) (float64, error) {
	if input == "" {
		return 0, errors.New("empty input string")
	}

	parts := strings.Fields(input)
	stack := []float64{}

	for _, part := range parts {
		if s, err := strconv.ParseFloat(part, 64); err == nil {
			stack = append(stack, s)
		} else {
			if len(stack) < 2 {
				return 0, errors.New("insufficient operands for operator " + part)
			}
			operand2 := stack[len(stack)-1]
			operand1 := stack[len(stack)-2]
			stack = stack[:len(stack)-2] // Pop two operands

			switch part {
			case "+":
				stack = append(stack, operand1+operand2)
			case "-":
				stack = append(stack, operand1-operand2)
			case "*":
				stack = append(stack, operand1*operand2)
			case "/":
				if operand2 == 0 {
					return 0, errors.New("division by zero")
				}
				stack = append(stack, operand1/operand2)
			default:
				return 0, errors.New("unknown operator: " + part)
			}
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid RPN expression: " + input)
	}

	return stack[0], nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("RPN Calculator (type 'q' to quit)")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if input == "q" {
			break
		}

		if input == "" {
			continue
		}

		result, err := CalculateRPN(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Println(result)
	}
}
