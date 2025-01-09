package main

import (
	"fmt"

	quote "rsc.io/quote/v4"
)

func main() {
	fmt.Println("Hello World!")

	/* 
	  Go is a function declared in "rsc.io/quote/v4" package
	*/
	fmt.Println(quote.Go())
}