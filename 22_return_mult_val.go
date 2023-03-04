package main
import "fmt"

func getFullName() (string, string, string) {
	return "test", "doli", "dani"
}
func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)

	// ignore salah 1 variable
	getName, _, _ := getFullName()
	fmt.Println(getName)
}