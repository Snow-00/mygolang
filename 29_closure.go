package main
import "fmt"

// closure adl kondisi di mana data berinteraksi hanya dg sekitarnya
// intinya kek local var engga bs dipake di func lain
func main() {
	counter := 0
	increment := func() {
		counter := 1 // ini biar counter di luar engga brubah nilainya
		fmt.Println("increment")
		counter++ // counter di luar jd brubah nilainya
		fmt.Println(counter)
	}

	increment()
	increment()
	fmt.Println(counter)
}