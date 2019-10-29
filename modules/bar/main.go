package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(echo())
	fmt.Println(quote.Hello())
}

func echo() string {
	return "bar"
}
