package main
import "fmt"

// membuat alias untuk func as param menggunakan command type (type declaration)
// contoh: alias untuk func(string) string
type Filter func(string) string

//func sayHelloWithFilter(name string, filter func(string) string)
func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("eko", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)
}