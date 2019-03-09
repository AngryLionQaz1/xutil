package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

var inputReader *bufio.Reader
var input2 string
var err error

func main() {

	t3()

}

func t3() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please enter your name:")
	input, _ = reader.ReadString('\n')

	switch input {

	case "A\r\n":
		fmt.Println("A")
	case "B\r\n":
		fmt.Println("B")
	case "C\r\n":
		fmt.Println("C")
	default:
		fmt.Println("X")

	}

}

func t2() {

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input:")
	input2, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was:%s\n", input2)
	}

}

func t() {
	fmt.Scanln(&firstName, &lastName)
	fmt.Println(firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read: ", f, i, s)
}
