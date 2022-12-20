package main

import "fmt"
        
        
func Merhaba(isim string)(cikti string){
    cikti ="Selam "+ isim

    return
}

func main(){
    selamla:= Merhaba("Melike")
    fmt.Println(selamla)
}

//fonksiyonlaırmızı test edeceğimiz için baş harflerini büyük yazmayı unutmuyoruz
