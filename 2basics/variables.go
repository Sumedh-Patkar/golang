package main

import "fmt"

func main() {

    // var declares 1 or more variables.
    fmt.Println("\nvar declares 1 or more variables.")
    var a = "initial"
    fmt.Println(a)

    // You can declare multiple variables at once.
    fmt.Println("\n\nDeclare multiple variables at once.")
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go will infer the type of initialized variables.
    fmt.Println("\n\nGo will infer the type of initialized variables.")
    var d = true
    fmt.Println(d)

    // Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
    fmt.Println("\n\nzero-valued variables")
    var e int
    fmt.Println(e)

    // The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case.
    fmt.Println("\n\n:= shorthand for declaring and initializing a variable.")
    f := "apple"
    fmt.Println(f)
}