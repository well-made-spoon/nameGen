package nameGen

import (
	"fmt"
	"math/rand"
	"time"
)

func RandomName() string {
	rand.Seed(time.Now().UnixNano())
	r1 := rand.Intn(len(Adjectives))
	r2 := rand.Intn(len(Nouns))
	str := fmt.Sprintf("%s %s", Adjectives[r1], Nouns[r2])
	return str
}
