package main

import (
	"fmt"
	"time"
)

var opMap map[string]func(int, int) int

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func initOpMap() {
	opMap = make(map[string]func(int, int) int)
	opMap["+"] = add
	opMap["-"] = sub
	opMap["*"] = mul
	opMap["/"] = div
}

func calculate(op string, a, b int) int {
	if v, ok := opMap[op]; ok {
		return v(a, b)
	}
	return 0
}

func main() {
	startTime := time.Now()
	initOpMap()
	test()
	elapsedTime := time.Since(startTime)
	fmt.Printf("runtime:%d \n", elapsedTime.Nanoseconds())
}

func assertEqual(testcase string, expected, o interface{}) bool {
	if expected != o {
		fmt.Printf("%s Failed! expected:%d output:%d\n", testcase, expected, o)
		return false
	}
	return true
}

func testCalculate(testcase, op string, a, b, expected int) bool {
	o := calculate(op, a, b)
	return assertEqual(testcase, expected, o)
}

func test() {
	if !testCalculate("test1", "+", 1, 2, 3) {
		return
	}
	if !testCalculate("test2", "+", 1, -2, -1) {
		return
	}

	if !testCalculate("test3", "-", 1, -2, 3) {
		return
	}
	if !testCalculate("test4", "-", 1, 2, -1) {
		return
	}

	if !testCalculate("test5", "*", 1, 2, 2) {
		return
	}
	if !testCalculate("test6", "*", 3, 7, 21) {
		return
	}
	if !testCalculate("test7", "*", -2, 5, -10) {
		return
	}
	if !testCalculate("test8", "*", -3, -1, 3) {
		return
	}
	if !testCalculate("test9", "*", 3, -4, -12) {
		return
	}

	if !testCalculate("test10", "/", 3, 1, 3) {
		return
	}
	if !testCalculate("test11", "/", 10, 2, 5) {
		return
	}
	if !testCalculate("test12", "/", 7, 2, 3) {
		return
	}
	if !testCalculate("test13", "/", -5, 3, -1) {
		return
	}
	if !testCalculate("test14", "/", 11, 11, 1) {
		return
	}

	fmt.Println("Success!")
}
