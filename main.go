package main

import (
	"fmt"
	"github.com/cokeys90/auto-bot-bithumb/application"
)

func main() {
	fmt.Println("Hello, World!")

	server := application.NewApplication()
	server.Listen()
}
