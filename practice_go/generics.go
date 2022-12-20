package main

import "fmt"


func main(){
    var sayi1 int =5
    fmt.Println(arttir(sayi1))

    
}
/*
func arttir(sayi int) int {
    return sayi + 1
}*/

//bu durumda arttir fonksiyonu sadee tam sayılar için kullanılabilecek dolayısıyla float değerler için kullanılamayacak, bu noktada generic aslında fonksiyon imzasında alınan değerin tipini birden fazla tipte olabilmesine olanak tanıyor

//buna göre fonksiyonumuzu tekrar yazalım

func arttir[n int | float64](sayi n) n {
    return sayi +1
}
