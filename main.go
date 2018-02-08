package main

import "fmt"

func main() {
	fmt.Println(returnOne())
}

//int before { = the type of variable being returned
func returnSaul() string {
	return "Saul"
}

func returnOne() int {
	return 1
}

func onePlusOne() int {
	two := 1 + 1
	return two
}

func sayHello() string {
	greeting := "Hi"
	fmt.Println(greeting)
	return "Hello!"
}

//name = the variable name. Here, name is a parameter
//string in () = the type of variable of `name`
//CHALLENGE: what type is being returned here?
func sayHelloTo(name string) string {
	//CHALLENGE: Print out the string "Hello, Pal"

	helloString := "Hello, " + name
	return helloString
}

func multiplythreebyfour() int {
	return 12
}
