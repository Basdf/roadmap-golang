package basics

import "fmt"

func StringIteration() {
	for i, ch := range "日本語" {
		fmt.Printf("%#U starts at byte position %d\n", ch, i)
	}
}
func MapIteration() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
func ChannelIteration() {
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)

	}
}
