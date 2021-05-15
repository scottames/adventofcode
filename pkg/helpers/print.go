package helpers

import "fmt"

const (
	oneInt          = 1
	printPartString = "--- Part %d ---\n\n"
	twoInt          = 2
)

func PrintPart1() {
	fmt.Printf(printPartString, oneInt)
}

func PrintPart2() {
	fmt.Printf(printPartString, twoInt)
}
