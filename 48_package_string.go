package main
import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("eko test", "eko"))
	fmt.Println(strings.Split("eko test testing", " "))
	fmt.Println(strings.ToUpper("Eko Test testing"))
	fmt.Println(strings.ToLower("Eko Test testing"))
	fmt.Println(strings.ToTitle("Eko Test testing"))  // sama kek uppercase
	fmt.Println(strings.Trim("     Eko Test testing     ", " "))
	fmt.Println(strings.ReplaceAll("Eko Test testing", "Eko", "budi"))
}