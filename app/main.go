package main

import (
	funcs "example/project/funcs"
	"fmt"
)

func main() {
	fmt.Println(funcs.Name)
	iarr := []int{1, 2, 3, 76543, 3, 5, 5, 6, 76, 7, 4, 56, 24, 2345}
	sArr := funcs.IntToStrArr(iarr)
	fmt.Println(sArr)
}
