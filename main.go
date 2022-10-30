package main

import (
	"fmt"
	"discord/events"
)

func main() {
	fmt.Println("Hello, Modules!")
	events.MessageCreate()
	events.MessageSend()
}