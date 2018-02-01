package main

import "fmt"

func main() {
	fmt.Println(returnOne())
}

//int before { = the type of variable being returned
func returnOne() int {
	return 1
}

func onePlusOne() int {
	return 0
}

func sayHello() string {
	//fmt.Println("Hello!")
	return ""
}

//name = the variable name. Here, name is a parameter
//string in () = the type of variable of `name`
//CHALLENGE: what type is being returned here?
func sayHelloTo(name string) string {
	//CHALLENGE: Print out the string "Hello, Pal"

	helloString := "Hello, "
	return helloString
}
