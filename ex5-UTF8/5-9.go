﻿package main

import "fmt"

func main() {
Loop:
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue Loop
		}
		fmt.Println(i)
	}
}
