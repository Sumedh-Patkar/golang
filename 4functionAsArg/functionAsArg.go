package main

import (
	_ "errors"	//this is ignored 
	"fmt"
	_ "log"	
	_ "os"
)

func main() {
	display(func(a, b int) int {return a+b})	
}

// A function is passed as an argument 
func display(a func(a, b int) int){

	// the function gets called with 3, 4 as arguments
	res := a(3, 4)
	fmt.Println("result = ", res)
	
}