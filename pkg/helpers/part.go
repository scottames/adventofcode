package helpers

import "fmt"

const (
	one     = 1
	partStr = "\n--- Part %d ---\n"
	two     = 2
)

func PrintPart1() {
	fmt.Println(StringPart1())
}

func StringPart1() string {
	return fmt.Sprintf(partStr, one)
}

func PrintPart2() {
	fmt.Println(StringPart2())
}

func StringPart2() string {
	return fmt.Sprintf(partStr, two)
}
