package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channnel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
