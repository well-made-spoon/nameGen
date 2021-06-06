package main

import (
	"flag"
	"fmt"
  "github.com/fatih/color"
	"github.com/bengadbois/flippytext"
	"github.com/well-made-spoon/nameGen/RandomName"
)

var numNames = flag.Int("n", 1, "Number of random names to generate")
var flippy = flag.Bool("flippy", false, "Print random name(s) using Flippy Text")
var color_flag = flag.Bool("color", false, "Print random name(s) using color")

func main() {
	flag.Parse()
	for i := 0; i < *numNames; i++ {
		name := nameGen.RandomName()
		if *flippy {
      if *color_flag{
      color.Set(color.FgGreen)
    }
			flippytext.New().Write(name)
		} else {
      if *color_flag{
      color.Set(color.FgGreen)
      }
			fmt.Println(nameGen.RandomName())
		}
	}
}
