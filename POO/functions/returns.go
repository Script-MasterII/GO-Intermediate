package main

import "fmt"

func sum(a, b int) int {
	return a + b
}
func sum2(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total

}
func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

func getValues(x int) (int, int, int) {
	return x, 2 * x, 3 * x
}
func getValues2(x int) (double int, triple int, quad int) {
	double = x * 2
	triple = x * 3
	quad = x * 4
	return
}
func main() {
	fmt.Println(sum(5, 54))
	fmt.Println(sum2(5, 5, 4, 4, 4, 4))
	printNames("Diego", "Carlos", "Enrique")
	fmt.Println(getValues(2))
	fmt.Println(getValues2(2))
}
