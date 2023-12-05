package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	var max int = 0
	for i, x := range os.Args {
		if i == 0 {
			continue
		}
		y, _ := strconv.Atoi(x)
		if y > max {
			max = y
		}
	}
	fmt.Println("Max:", max)
}
