package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("The Larger Testing Unit", func() {
	It("should do something", func() {
		Expect(returnOne()).To(Equal(1))
	})
})
