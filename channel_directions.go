package main

import "fmt"

func ping(pings chan<- string, msg string) {
	msg := <-pings
	pongs <- msg
}
