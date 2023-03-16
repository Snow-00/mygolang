package main
import "fmt"

/**
interface adl tipe data abstract
interface berisi definisi2 method
*/

type HasName interface {
  getName() string
}

func sayHello(hasName HasName) {
  fmt.Println("hello", hasName.getName())
}

type Person struct {
  Name string
}

// func harus sama persis confignya sesuai yg ada di interface
func (person Person) getName() string {
  return person.Name
}

type Animal struct {
  Name string
}

func (animal Animal) getName() string {
  return animal.Name
}

func main() {
  var eko Person
  eko.Name = "eko"
  sayHello(eko)

  cat := Animal {
    Name: "puss",
  }
  sayHello(cat)
}