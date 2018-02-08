package main

// These are libraries. They contain code that someone else wrote that we can use
// Why might libraries be useful?
import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Work with return values and basic functions", func() {

	It("should compare a variable and a hardcoded value", func() {
		one := 1
		//This is an expect statement. We compare the variable "one" on the left to the actual integer 1 on the right
		Expect(one).To(Equal(1))

		// CHALLENGE: WHAT DO YOU THINK WILL HAPPEN IF WE CHANGE THE 1 TO A 2?
	})

	It("should compare a variable and the return value of a function", func() {
		one := 1
		// Here, we are calling the function returnOne, defined in `main.go`
		returnedValue := returnOne()

		Expect(returnedValue).To(Equal(one))
	})

	It("should add 1 + 1 together", func() {
		answer := onePlusOne()

		// Let's go find the function `onePlusOne` in `main.go` and create the right return value
		Expect(answer).To(Equal(2))
	})

	It("should be able to write out text", func() {
		helloString := sayHello()

		// Let's go find the function `sayHello` in `main.go` and create the right return value
		Expect(helloString).To(Equal("Hello!"))
	})
})

var _ = Describe("Work with parameters and troubleshooting functions", func() {

	It("should be able to write out `Hello, Pal!`", func() {
		// Here, "Pal" is called a parameter. It is a piece of data being send into
		// the function so the function can use that data
		helloString := sayHelloTo("Pal")
		helloSaul := sayHelloTo("Saul")

		// Let's go find the function `sayHelloTo` in `main.go` and create the right return value
		Expect(helloString).To(Equal("Hello, Pal"))
		Expect(helloSaul).To(Equal("Hello, Saul"))
	})

	It("should have a function called multiplyThreeTimesFour", func() {
		//CHALLENGE: Oh no! What are we seeing in the output???
		var returnedNumber int
		returnedNumber = multiplythreebyfour()

		Expect(returnedNumber).To(Equal(12))
	})

})
