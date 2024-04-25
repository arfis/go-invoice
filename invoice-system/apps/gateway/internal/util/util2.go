package util

import "fmt"

type TestingFunc2 func(string, int)

// ServeHTTP calls f(w, r).
func (f TestingFunc2) Test(s string, i int) {
	f(s, i)
}

func testFunc(s string, i int) {
	fmt.Println("Testing")
}

func testingTypes() {
	var testingFuncInstance = TestingFunc2(testFunc)
	testingFuncInstance.Test("string", 12)
}
