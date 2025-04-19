package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	ReadFile()
}

func ReadFile() {
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
