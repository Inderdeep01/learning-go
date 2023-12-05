package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Create("data")
	if err != nil {
		log.Fatal(err)
	}
	iPrime := []int{2, 3, 5, 7, 11, 13}
	var sPrime []string
	for _, x := range iPrime {
		sPrime = append(sPrime, strconv.Itoa(x))
	}
	for _, x := range sPrime {
		_, err := file.WriteString(x + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	file.Close()
	fmt.Println("Data written to the file")
	file, err = os.Open("data")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Read error", err)
	}
	_, err = os.Stat("data")
	fmt.Println("FILE does not exists :", os.IsNotExist(err))
	if os.IsNotExist(err) {
		fmt.Println("error")
		fmt.Println(os.ErrNotExist)
	} else {
		file, err := os.OpenFile("data", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		if _, err := file.WriteString("19"); err != nil {
			log.Fatal(err)
		}
	}

}
