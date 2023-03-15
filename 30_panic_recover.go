package main
import "fmt"

// panic itu lgsg nampilin error sesuai yg kt isi dan menghentikan apps
// nil adl value yg dimunculkan recover ketika tidak ada error panic
// recover ditarok di defer func
func endApp() {
  message := recover()
  if message != nil {
    fmt.Println("error dengan message", message)
  }
  fmt.Println("aplikasi slesai")
}

func runApp(error bool) {
  defer endApp()
  if error {
    panic("aplikasi error")
  }
  fmt.Println("aplikasi berjalan")
}

func main() {
  runApp(true)
  fmt.Println("test")  // code ini jd ttp jalan krn ada recover
}