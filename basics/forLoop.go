package basics

import "fmt"

func ThreeComponentLoop() {
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)
}
func WhileLoop() {
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)
}

func InfiniteLoop() {
	sum := 0
	for {
		sum++ // repeated forever
		if sum == 10 {
			break
		}
	}
	fmt.Println(sum)
}
func RangeLoop() {
	strings := []string{"hello", "world"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
}
func ExitLoop() {
	sum := 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		sum += i
	}
	fmt.Println(sum)
}
