package chapter_3

import (
	"fmt"
	"math/rand"
)

//LetsLoop loops with n as a condition
func LetsLoop(n int) {
	for {
		if rand.Intn(n+1) == n {
			for i := 0; i < 3; i++ {
				fmt.Printf("Hey Guys i found my number %v\n", n)
			}
			break
		}
	}
}
