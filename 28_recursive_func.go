package main
import "fmt"

// stack overflow = infinite loop

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	loop:= factorialLoop(10)
	fmt.Println(loop)

	recursive := factorialRecursive(10)
	fmt.Println(recursive)
}