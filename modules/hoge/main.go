package main

import (
	"fmt"
	"rsc.io/quote"
)

//const Hoge = 1

func main() {
	fmt.Println(echo())
	fmt.Println(quote.Hello())
}

func echo() string {
	return "hoge"
}
