package basics

import (
	"fmt"
	"time"
)

func Conditionals() {
	x := 10
	if x > 5 {
		fmt.Println("x es mayor que 5")
	} else if x < 5 {
		fmt.Println("x es menor que 5")
	} else {
		fmt.Println("x es igual a 5")
	}
	if y := 10; y > 5 {
		fmt.Println("y es mayor que 5")
	}
}

func Switch() {
	today := time.Now()
	switch today.Day() {
	case 1, 2, 3, 4:
		fmt.Println("It's a weekday.")
		fallthrough // hace que el switch entre al siguiente caso aunque no cumpla la condicion (solo se puede usar al final de un caso)
	case 5, 10, 15:
		fmt.Println("Clean your house.")
	case 25, 26, 27:
		fmt.Println("Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
	fmt.Println("switch initializer Statement")
	switch today := time.Now(); {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some wine.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}
}
