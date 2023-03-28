/** 
nama package hrs sama dengan dir di atas file ini
dlm 1 package ga blh ada func / struct dg nama yg sama, kalo udh beda 
package itu br blh
*/

package helper
import "fmt"

// kalo utk package var nya hrs cara gini, kalo cara := bakal error
var version = 1
var Application = "belajar golang"

// nama func hrs diawali huruf besar, kalo ga error nanti
func SayHello(name string) {
  fmt.Println("hello", name)
}

func sayGoodbye(name string) {
  fmt.Println("Goodbye", name)
}