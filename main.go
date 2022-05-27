package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	val1, val2 := getValues(reader)
	operation := getOperation(reader)

	var result float64

	switch operation {
	case '+':
		result = add(val1, val2)
	case '-':
		result = subtract(val1, val2)
	case '*':
		result = multiply(val1, val2)
	case '/':
		result = divide(val1, val2)
	default:
		panic("Invalid operation!")
	}

	fmt.Printf("\n%.2f %c %.2f = %.2f\n", val1, operation, val2, result)
}

func getValues(r *bufio.Reader) (float64, float64) {
	fmt.Print("Value 1: ")
	input1, _ := r.ReadString('\n')
	num1, err1 := strconv.ParseFloat(strings.TrimSpace(input1), 64)

	if err1 != nil {
		panic("Invalid input.")
	}

	fmt.Print("Value 2: ")
	input2, _ := r.ReadString('\n')
	num2, err2 := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	if err2 != nil {
		panic("Invalid input.")
	}

	return num1, num2
}

func getOperation(r *bufio.Reader) byte {
	fmt.Print("Select an operation (+ - * /): ")
	input, err := r.ReadByte()

	if err != nil || input < 42 || input > 47 || input == 44 || input == 46 {
		panic("Invalid input!")
	}

	return input
}

func add(val1, val2 float64) float64 {
	return val1 + val2
}

func subtract(val1, val2 float64) float64 {
	return val1 - val2
}

func multiply(val1, val2 float64) float64 {
	return val1 * val2
}

func divide(val1, val2 float64) float64 {
	return val1 / val2
}
