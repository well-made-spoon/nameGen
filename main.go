package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/bengadbois/flippytext"
)

func main() {
	times := "1"
	if len(os.Args) > 1 {
		times = os.Args[1]
	}
	timesInt, _ := strconv.Atoi(times)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < timesInt; i++ {
		r1 := rand.Intn(len(Adjectives))
		r2 := rand.Intn(len(Nouns))
		str := fmt.Sprintf("%s %s", Adjectives[r1], Nouns[r2])
		flippytext.New().Write(str)
	}
}
