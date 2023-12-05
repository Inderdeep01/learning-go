package funcs

import "strconv"

const Name string = "Cosntant String"

var Name2 string = "Variable string"

func IntToStrArr(arr []int) []string {
	var res []string
	for _, x := range arr {
		res = append(res, strconv.Itoa(x))
	}
	return res
}
