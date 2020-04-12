package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	ADD = "addition"
	SUB = "subtraction"
	MUL = "multiplication"
	DIV = "division"
)

var (
	userOperator string
)

func main() {
	fmt.Println("Welcome to my simple calculator app. What name can i address you with")
	var input string
	fmt.Scanln(&input)

	fmt.Println("Thanks", input, ". Enter operation symbol: [a] Additon [s] Subtraction [m] Multiplication [d] Division")
	var operation string
	fmt.Scanln(&operation)

	operationResponse := verifyOperationResponse(operation)
	if operationResponse == "error" {
		fmt.Println(input, "the operation", operation, "you entered isnt a valid one. quitting operation...")
		os.Exit(2)
	}
	userOperator = operationResponse
	val1, val2 := getValues()
	intVal1, intVal2 := verifyUserValues(val1, val2)
	performOperation(operationResponse, intVal1, intVal2)
}

func verifyOperationResponse(response string) string {
	var operation string
	if response == "a" {
		operation = "addition"
	} else if response == "s" {
		operation = "subtraction"
	} else if response == "m" {
		operation = "multiplication"
	} else if response == "d" {
		operation = "division"
	} else {
		operation = "error"
	}
	return operation
}

func getValues() (string, string) {
	fmt.Println("Enter first value")
	var value1 string
	fmt.Scanln(&value1)
	fmt.Println("Enter second value")
	var value2 string
	fmt.Scanln(&value2)
	return value1, value2
}

func verifyUserValues(v1 string, v2 string) (int, int) {
	getint1, err1 := strconv.Atoi(v1)
	if err1 != nil {
		fmt.Println("It doesnt look like", v1, "is a number. Exiting...")
		os.Exit(2)
	}
	getint2, err2 := strconv.Atoi(v2)
	if err2 != nil {
		fmt.Println("It doesnt look like", v2, "is a number. Exiting...")
		os.Exit(2)
	}
	return getint1, getint2
}

func performOperation(op string, v1 int, v2 int) {
	var result int
	var resultFloat float64
	if op == ADD {
		result = v1 + v2
		fmt.Println("Result of the", userOperator, "is", result)
	} else if op == SUB {
		result = v1 - v2
		fmt.Println("Result of the", userOperator, "is", result)
	} else if op == MUL {
		result = v1 * v2
		fmt.Println(result)
	} else if op == DIV {
		if v2 == 0 {
			fmt.Println("Division by zero is not allowed")
			os.Exit(2)
		}
		resultFloat = float64(v1) / float64(v2)
		fmt.Println("Result of the", userOperator, "is", resultFloat)
	} else {
		fmt.Println("Undefined operation")
		os.Exit(2)
	}
}
