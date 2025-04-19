package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("in main")
	ReadFile()
	fmt.Println(" going out of main")
}

func ReadFile() {
	fmt.Println("in read file ")
	s, err := callerror()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(s)
}

func callerror() (string, error) {
	return "", errors.New("hello raj this is error")
}
