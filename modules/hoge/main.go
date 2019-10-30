package main

import (
	"fmt"
	"rsc.io/quote"
)

// This is lint error!
const Hoge = 1

func main() {
	// This is build error!
	// ffmt.Println(echo())
	fmt.Println(echo())
	fmt.Println(quote.Hello())
}

func echo() string {
	// This is test error!
	// return "foo"
	return "hoge"
}
