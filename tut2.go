package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("What is your age? ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	trimmedInput := strings.TrimSpace(input)
	age, err := strconv.ParseInt(trimmedInput, 10, 32)
	fmt.Println("Your age is", age)
	if err == nil {
		if age < 5 {
			fmt.Println("Too young for school")
		} else if age == 5 {
			fmt.Println("Go to kindergarten")
		} else if age > 5 && age <= 17 {
			fmt.Printf("Go to grade %d \n", age-5)
		} else {
			fmt.Println("Go to College")
		}
	}
}
