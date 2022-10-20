package main

import (
	"fmt"
	"os"
	"strings"
)

func printUsage() {
	fmt.Println("Invalid usage! Usage:\n", os.Args[0], "<username>", "<message>")
}

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		printUsage()
		os.Exit(1)
	}

	uname := args[0]
	messages := strings.Join(args[1:], " ")

	fmt.Println("Sending messages to", uname)
	tok, err := Sent(&uname, &messages)
	if err != nil {
		fmt.Println("Couldn't sent to", uname, "because", err.Error())
		os.Exit(1)
	}

	fmt.Println("Success, the device token:", *tok)
}
