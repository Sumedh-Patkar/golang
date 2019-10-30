package main

import (
	_ "errors"	//this is ignored 
	"fmt"
	_ "log"	
	_ "os"
)

type Student struct {
	name string
	age int
}

func (m *Student) getName() string { 
	return m.name 
}

func (m *Student) setName(newName string) {
	m.name = newName	
}

func main() {
	x := &Student{name:"sumedh", age:4}	

	fmt.Println(x.name)
	fmt.Println(x.age)
	fmt.Println(x.getName())
	x.setName("patkar")
	fmt.Println(x.getName())	
}