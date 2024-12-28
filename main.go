package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	t := mustToken()

	fmt.Println(t)
}

func mustToken() string {
	token := flag.String("token", "", "telegram bot token")

	flag.Parse()

	if *token == "" {
		log.Fatal("telegram bot token required")
	}

	return *token
}
