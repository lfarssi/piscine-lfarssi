package main

import (
	"os"
	"math"
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
		if isOverflow(val1, val2, "+") {
			return // Overflow detected
		}
		result = add(val1, val2)
	case "-":
		if isOverflow(val1, val2, "-") {
			return // Overflow detected
		}
		result = sub(val1, val2)
	case "*":
		if isOverflow(val1, val2, "*") {
			return // Overflow detected
		}
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

// isOverflow checks if the operation causes an overflow.
func isOverflow(num1, num2 int, op string) bool {
	switch op {
	case "+":
		return (num2 > 0 && num1 > math.MaxInt-int(num2)) || (num2 < 0 && num1 < math.MinInt-int(num2))
	case "-":
		return (num2 < 0 && num1 > math.MaxInt+int(num2)) || (num2 > 0 && num1 < math.MinInt+int(num2))
	case "*":
		if num1 > 0 && num2 > 0 {
			return num1 > math.MaxInt/num2
		}
		if num1 < 0 && num2 < 0 {
			return num1 < math.MaxInt/num2
		}
		if num1 > 0 && num2 < 0 {
			return num2 < math.MinInt/num1
		}
		if num1 < 0 && num2 > 0 {
			return num1 < math.MinInt/num2
		}
		return false
	}
	return false
}

// printResult converts an integer to a string and writes it to stdout with a newline.
func printResult(result int) {
	if result == 0 {
		os.Stdout.Write([]byte("0\n"))
		return
	}

	if result < 0 {
		os.Stdout.Write([]byte("-"))
		result = -result
	}

	// Build the string representation of the number
	var digits []byte
	for result > 0 {
		digit := result % 10
		digits = append([]byte{byte('0' + digit)}, digits...)
		result /= 10
	}

	os.Stdout.Write(digits)
	os.Stdout.Write([]byte("\n"))
}

// printError prints an error message followed by a newline to stdout.
func printError(message string) {
	os.Stdout.Write([]byte(message + "\n"))
}
