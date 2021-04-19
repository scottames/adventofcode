package helpers

// Absolute - returns the absolute of an integer
func Absolute(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Add - returns the sum of two integers
func AddInt(x int, y int) int {
	return x + y
}

// Divide - returns the division of two integers
func DivideInt(x int, y int) int {
	return x / y
}

// Multiply - returns the product of two integers
func MultiplyInt(x int, y int) int {
	return x * y
}

// Subtract - returns the difference of two integers
func SubtractInt(x int, y int) int {
	return x - y
}
