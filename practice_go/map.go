package main
import "fmt"

  type insan struct {
   kisi1, kisi2, kisi3 string   
  }
  
  func main(){
      var m map[string] insan
      m = make(map[string]insan)
      m["isim"] = insan {
            "A","M","N",
      }
      fmt.Println(m["isim"])
  }