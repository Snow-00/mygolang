package main
import "fmt"

// defer digunakan utk menjalankan func b walopun func a gagal
func logging() {
  fmt.Println("selesai panggil func")
}

func runApp(value int) {
  defer logging()    // taro defer di atas, kalo di bawah sama aj ga jl
  fmt.Println("run app")
  result := 10 / value
  fmt.Println("result", result)
}

func main() {
  runApp(0)
}