package main

import (
	"bufio"
	f "fmt"
	"log"
	"os"
)

var print = f.Println

func main() {
	print("Hello World in GO")
	print("Enter your Name:")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		print("Hello,", name, " form GO")
	} else {
		log.Fatal(err)
	}
}
