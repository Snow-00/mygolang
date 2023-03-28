/**
di golang bs membuat func yg bakal langsung dirun ketika package nya
diakses, nama func nya init (wajib init)
func init cocok dipake saat package nya buat komunikasi dg database
(func init nya buat membuka koneksi ke database)
*/

package database
import "fmt"

var connection string

func init() {
  connection = "MySQL"
  fmt.Println("init dipanggil")
}

func GetDatabase() string {
  return connection
}