package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func Add(left, right float64) float64 {
	return left + right
}

func Subtract(left, right float64) float64 {
	return left - right
}

func Multiply(left, right float64) float64 {
	return left * right
}

func Divide(left, right float64) float64 {
	return left / right
}

func ParseOperands(args []string) (string, float64, float64, error) {
	var leftOperand, rightOperand float64
	var err error
	if len(args) != 3 {
		err = errors.New("Usage: ./main <operation> <operand1> <operand2>")
		return "", 0, 0, err
	}
	operation := args[0]
	leftOperand, err = strconv.ParseFloat(args[1], 64)
	if err != nil {
		return "", 0, 0, err
	}
	rightOperand, err = strconv.ParseFloat(args[2], 64)
	if err != nil {
		return "", 0, 0, err
	}
	return operation, leftOperand, rightOperand, nil

}

func MakeOperation(operation string, l, r float64) (float64, error) {
	switch operation {
	case "+":
		return Add(l, r), nil
	case "-":
		return Subtract(l, r), nil
	case "*":
		return Multiply(l, r), nil
	case "/":
		if r == 0 {
			err := errors.New("Zero division!")
			return 0, err
		}
		return Divide(l, r), nil
	}
	err := errors.New("Unknown operation!")
	return 0, err
}

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Usage: ./main <operation> <operand1> <operand2>")
		os.Exit(1)
	}

	operation, leftOperand, rightOperand, err := ParseOperands(os.Args[1:])
	if err != nil {
		fmt.Println("Incorrect input data. Err:", err)
		os.Exit(1)
	}
	res, err := MakeOperation(operation, leftOperand, rightOperand)
	if err != nil {
		fmt.Println("Error in process of calculating result, err = ", err)
		os.Exit(1)
	}
	fmt.Println("Result = ", res)
}
