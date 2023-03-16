package main
import "fmt"

// coding lanjut dr code 32
type Customer struct {
  Name, Address string
  Age           int
}

// func biasa
//func sayHi(customer Customer, name string) {
//  fmt.Println("hi", name, "my name is", customer.Name)
//}

// struct func / struct method
func (customer Customer) sayHi(name string) {
  fmt.Println("hi", name, "my name is", customer.Name)
}

func (test Customer) sayHuu() {
  fmt.Println("huu from", test.Name)
}

func main() {
  var eko Customer
  eko.Name = "eko"
  eko.Address = "indonesia"
  eko.Age = 29
  
  //sayHi(eko, "joko")
  eko.sayHi("jokoo")
  eko.sayHuu()
}