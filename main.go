package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("in main")
	WriteFile()
	ReadFile()
	fmt.Println(" going out of main")
}

func WriteFile() {
	fmt.Println("in write file")
	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	_, err = f.WriteString("hello world")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("file written successfully")
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
	return "", nil
}
