package main
import "fmt"

func getFullName2() (firstName, middleName, lastName string) {
  firstName = "Jeffri"
  middleName = "Doni"
  lastName = "baghdad"

  return    // sudah tidak perlu sebutin 1" nilai yg ingin direturn
}

func main() {
  firstName, middleName, lastName := getFullName2()
  fmt.Println(firstName, middleName, lastName)
}