package main

import (
	"errors"
	"fmt"
	"log"	
	"os"
)

func main() {
	fmt.Println("Simple Calculator")
	a, b := 6, 3

	fmt.Println("Add:\t", add(a, b))
	fmt.Println("Sub:\t", sub(a, b))
	fmt.Println("Mul:\t", mul(a, b))

	result, err := div(a, b)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Println("Div:\t", result)	
	
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}


func div(a, b int) (int, error) {
	// error handling is done here
	if b == 0 {	
		return -1, errors.New("Divide by zero error")
	} 	
	return a/b, nil
}
