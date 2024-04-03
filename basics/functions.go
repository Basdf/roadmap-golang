package basics

import "fmt"

// A name must begin with a letter, and can have any number of additional letters and numbers.
// A function name cannot start with a number.
// A function name cannot contain spaces.
// If the functions with names that start with an uppercase letter will be exported to other packages. If the function name starts with a lowercase letter, it won't be exported to other packages, but you can call this function within the same package.
// If a name consists of multiple words, each word after the first should be capitalized like this: empName, EmpAddress, etc.
// function names are case-sensitive (car, Car and CAR are three different variables).
var (
	Area = func(l int, b int) int {
		return l * b
	}
)

func SayHello() {
	fmt.Println("Hello World!")
}
func Sum(a int, b int) int {
	return a + b
}

func Minus(a int, b int) (result int) {
	result = a - b
	return
}

func RectangleArea(base int, height int) (parameter, area int) {
	parameter = 2 * (base + height)
	area = base * height
	return
}

func Multiply(a int, b int) int {
	return a * b
}
func PartialMultiply(x int) func(int) int {
	return func(y int) int {
		return Multiply(x, y)
	}
}
func SquareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}