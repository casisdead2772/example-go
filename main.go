package main

import (
	"example-go/helper"
	"fmt"
)

var Global = 5

func main() {
	fmt.Println(Global)
	UseGlobal()
	fmt.Println(Global)
	helper.Help()
}

func UseGlobal() {
	defer func(copy int) {
		Global = copy
	}(Global)

	Global = 123
	fmt.Println(Global)
}
