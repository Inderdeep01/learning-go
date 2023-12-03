package main

import (
	"fmt"
	"strconv"
)

func main() {
	var v1 string = string(65)
	fmt.Println(v1)
	v2 := "654"
	v3, _ := strconv.Atoi(v2)
	fmt.Println(v3)
	v4 := "12.34"
	if v5, err := strconv.ParseFloat(v4, 64); err == nil {
		fmt.Println(v5)
	}
	v5 := fmt.Sprintf("%f", 31.5)
	fmt.Println(v5)
}
