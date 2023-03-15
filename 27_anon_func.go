package main
import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("ur blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

// bisa membuat func langsung dimasukin ke var / langsung jd param
func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("eko", blacklist)
	// terlalu capek kalo gini
	registerUser("root", func(name string) bool{
		return name == "root"
	})
}