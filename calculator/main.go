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

func Convert(args []string) (float64, float64, error) {
	var leftOperand, rightOperand float64
	var err error
	leftOperand, err = strconv.ParseFloat(args[0], 64)
	if err != nil {
		return 0, 0, err
	}
	rightOperand, err = strconv.ParseFloat(args[1], 64)
	if err != nil {
		return 0, 0, err
	}
	return leftOperand, rightOperand, nil

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
		fmt.Println("Incorrect amount of params!")
		os.Exit(1)
	}
	operation := os.Args[1]
	var leftOperand, rightOperand, res float64
	var err error
	leftOperand, rightOperand, err = Convert(os.Args[2:4])
	if err != nil {
		fmt.Println("Incorrect input data. Err:", err)
		os.Exit(1)
	}
	res, err = MakeOperation(operation, leftOperand, rightOperand)
	if err != nil {
		fmt.Println("Error in process of calculating result, err = ", err)
		os.Exit(1)
	}
	fmt.Println("Result = ", res)
}
