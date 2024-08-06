package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return // Invalid number of arguments
	}

	val1Str := os.Args[1]
	op := os.Args[2]
	val2Str := os.Args[3]

	val1, err1 := parseInt(val1Str)
	val2, err2 := parseInt(val2Str)

	if err1 != nil || err2 != nil {
		return // Invalid values
	}

	var result int
	switch op {
	case "+":
		result = add(val1, val2)
	case "-":
		result = sub(val1, val2)
	case "*":
		result = mult(val1, val2)
	case "/":
		if val2 == 0 {
			printError("No division by 0")
			return
		}
		result = div(val1, val2)
	case "%":
		if val2 == 0 {
			printError("No modulo by 0")
			return
		}
		result = mod(val1, val2)
	default:
		return // Invalid operator
	}

	printResult(result)
}

// parseInt converts a string to an integer manually without using strconv.
func parseInt(s string) (int, error) {
	if len(s) == 0 {
		return 0, &parseIntError{s}
	}

	var num int
	sign := 1
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			return 0, &parseIntError{s}
		}
		num = num*10 + int(ch-'0')
	}

	return num * sign, nil
}

// parseIntError represents an error for parseInt function.
type parseIntError struct {
	str string
}

func (e *parseIntError) Error() string {
	return "invalid syntax for parsing integer: " + e.str
}

func add(num1, num2 int) int {
	return num1 + num2
}

func sub(num1, num2 int) int {
	return num1 - num2
}

func mult(num1, num2 int) int {
	return num1 * num2
}

func div(num1, num2 int) int {
	return num1 / num2
}

func mod(num1, num2 int) int {
	return num1 % num2
}

// printResult converts an integer to a string and writes it to stdout manually.
func printResult(result int) {
	if result == 0 {
		os.Stdout.Write([]byte("0"))
		return
	}

	if result < 0 {
		os.Stdout.Write([]byte("-"))
		result = -result
	}

	// Build the string representation of the number
	var temp []byte
	for result > 0 {
		digit := result % 10
		temp = append([]byte{byte('0' + digit)}, temp...)
		result /= 10
	}

	os.Stdout.Write(temp)
}

// printError prints an error message to stdout.
func printError(message string) {
	os.Stdout.Write([]byte(message))
}
