# Basics of Go

In this section, we cover basics of packages, functions and variables

## Packages

  The program corresponding to this section is [packages.go](https://github.com/Sumedh-Patkar/golang/blob/master/2basics/packages.go)

  ### What are they
  
  Every Go program is made up of packages.

  Programs start running in `package main`.
  By convention, the package name is the same as the last element of the import path. 
  For instance, the `"math/rand"` package comprises files that begin with the statement `package rand`. 

  ### Imports
  
  In [packages.go](https://github.com/Sumedh-Patkar/golang/blob/master/2basics/packages.go), we have imported packages 
  `"fmt"`, `"math"` and `"math/rand"` using the `import` statement
  
  You can also write multiple import statements, like: 
  ```Go
  import "fmt"
  import "math"
  ```

  But it is good style to use the factored import statement. 
  
  ```Go
  import (
    "fmt"
    "math"
  )
  ```

  ### Exported Names

  In Go, a name is exported if it **begins with a _capital letter_**. 
  For example, `Pizza` is an exported name, as is `Pi`, which is exported from the `math` package. 
  
  A name beginning with a _small letter_ is not exported.
  
## Variables

  The program corresponding to this section is [variables.go](https://github.com/Sumedh-Patkar/golang/blob/master/2basics/variables.go)


  A Variable can be created in two ways
  
  1. **Using the `var` keyword:**

  ```Go
  var i int
  i = 4
  ```
    
  Variables with initializers
  
  A var declaration can include initializers, one per variable.
  If an initializer is present, the _type can be omitted_; the variable will take the type of the initializer. 

  ```Go
  var i, j int = 1, 2
  ```

  2. **Using short variable declaration:**

   Inside a function, the := short assignment statement can be used in place of a var declaration with **implicit type**. 

  ```Go
  x := "Hello World"
  ```
  Notice the : before the = and that no type was specified. 
  
  The type is not necessary because the Go compiler is able to infer the type based on the literal value you assign the variable.  
  
  
  *Note:* Variables declared without an explicit initial value are given their zero value.
  The zero value is:
   * 0 for numeric types,
   * false for the boolean type, and
   * "" (the empty string) for strings.


## Functions

  The program corresponding to this section is [functions.go](https://github.com/Sumedh-Patkar/golang/blob/master/2basics/functions.go)
  
  A function is an independent section of code that maps zero or more input parameters to zero or more output parameters. <br /> 
  Functions (also known as procedures or subroutines) are often represented as a black box: (the black box represents the function)

      
  <p align="center">
    <img src="./images/function-blackbox.png" alt="Function Blackbox" width="400">
  </p>

  Functions start with the keyword `func`, followed by the function's name. <br />
  The parameters (inputs) of the function are defined like this: `name type, name type, …`
  
  Example of a function 
  
  ```Go
  func square(a int) int {
	  return a * a
  }
  ```
  
  Excellent documentation for functions can be found [here](https://www.golang-book.com/books/intro/7)

