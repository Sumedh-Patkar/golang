package main

import "fmt"

/* 
Here’s a function that takes two ints 
and returns their sum as an int.
*/
func plus(a int, b int) int {

    return a + b
}

/*
When you have multiple consecutive parameters of the same type, 
you may omit the type name for the like-typed parameters 
up to the final parameter that declares the type.
*/
func plusPlus(a, b, c int) int {
    return a + b + c
}

/*
Function with multiple return values.
The (int, int) in this function signature shows that 
the function returns 2 ints.
*/
func vals() (int, int) {
    return 3, 7
}

/*
Variadic functions can be called with any number of trailing arguments. 
For example, fmt.Println is a common variadic function.
Here’s a function that will take an arbitrary number of ints as arguments.
*/
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println("Sum = ", total)
}

func main() {

	fmt.Println("\n#################################")
	fmt.Println("###### Simple functions ######")
	fmt.Println("#################################")
    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)

    // Here we use the 2 different return values from the call 
    // with multiple assignment.
    fmt.Println("\n##############################################")
    fmt.Println("### Function with multiple return values ###")
    fmt.Println("##############################################")
    a, b := vals()
    fmt.Println("a = ", a)
    fmt.Println("b = ", b)


    // Variadic functions
    // Can be called with any number of arguments
    fmt.Println("\n########################################################################")
    fmt.Println("### Variadic functions i.e. functions with any number of arguments ###")
    fmt.Println("########################################################################")
    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...)
}

