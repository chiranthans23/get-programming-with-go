package main

import (
	"fmt"

	c1 "github.com/get-programming-with-go/chapter_1"
	c2 "github.com/get-programming-with-go/chapter_2"
	c3 "github.com/get-programming-with-go/chapter_3"
)

func main() {
	c1.PrintMyName("chiranthan")
	c1.PrintNonEnglishString()

	c2.CalSpeedToMars()

	fmt.Printf("If you say Chiranthan is GOlang expert, I say %v\n", c3.ContainsString("Cisco", "Cisco"))
	c3.ComapreValues(4, 3, "<=")
	c3.LetsLoop(3)
}
