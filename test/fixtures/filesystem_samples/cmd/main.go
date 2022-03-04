/*
	Joao Daniel Rufino
	github.com/JoaoDanielRufino
*/

package main

import (
	"fmt"

	"github.com/JoaoDanielRufino/gcloc/test/fixtures/filesystem_samples/pkg/hello"
	"github.com/JoaoDanielRufino/gcloc/test/fixtures/filesystem_samples/pkg/sum"
)

// main function
func main() {
	hello.SayHello()

	a := 2
	b := 3
	// Sum a + b
	total := sum.Sum(a, b)

	fmt.Println(total)
}
