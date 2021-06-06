package main

import (
	"flag"
	"fmt"
	"github.com/bengadbois/flippytext"
	"github.com/well-made-spoon/nameGen"
)

var numNames = flag.Int("n", 1, "Number of random names to generate")
var flippy = flag.Bool("flippy", false, "Print random name(s) using Flippy Text")

func main() {
	flag.Parse()
	for i := 0; i < *numNames; i++ {
		name := nameGen.RandomName()
		if *flippy {
			flippytext.New().Write(name)
		} else {
			fmt.Println(nameGen.RandomName())
		}
	}
}
