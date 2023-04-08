/**
balikan dari pkg strconv ada 2, hasil dan error
*/

package main
import (
	"fmt"
	"strconv"
)

func main() {
	// parse utk conversion dr str ke tipe data lain
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}
	
	// angka, base angka misal hexa binary decimal okta, int_ (int32)
	number, err := strconv.ParseInt("1000000", 10, 64) 
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	// format itu ngubah tipe data lain jd str
	value := strconv.FormatInt(1000000, 16) // 10 = base angka decimal
	fmt.Println(value)

	// ada atoi (str to int) dan itoa (int to str)
	valueInt, _ := strconv.Atoi("2000200")
	fmt.Println(valueInt)
}