// name of the current package
package main

import (
	"fmt"
	"math"
	"sort"
)

//fmt 	Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
//math  Package math provides basic constants and mathematical functions. 
//sort  Package sort provides primitives for sorting slices and user-defined collections. 



func main() {
	fmt.Println("\nPackage math provides basic constants and mathematical functions.")
	x := math.Abs(-2)
	fmt.Printf("%.1f\n", x)

	fmt.Println("\nPackage sort provides primitives for sorting slices and user-defined collections. ")
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(s)
	fmt.Println(s)
}