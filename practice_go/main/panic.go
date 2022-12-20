package main

import "fmt"

func main(){
    sayilar := make([]int, 5)
    sayilar[6] = 10
    fmt.Println(len(sayilar)) // cevap 5 normalde
    //Ama gördüğün gibi burada olmayan inidise sayi atamaya calisiyoruz
    
}

//yukarida made fonksiyonu sayesinde sayilar adinda 5 birimden oluşan int bir dizi olıuturduk 


/*
func TamIsim(Ad *string, Soyad *string) {
    if Ad == nil {
        panic("Ad nil olamaz")
    }
    if Soyad == nil {
        panic("Soyad nil olamaz")
    }
    fmt.Printf("%s %s\n", *Ad, *Soyad)
    fmt.Println("TamIsim fonksiyonu bitti")
}
​
func main() {
    Ad := "Yusuf"
    TamIsim(&Ad, nil)
    fmt.Println("Ana fonksiyon da bitti")
}

////////////////////////////////


import (
    "fmt"
)

func TamIsim(Ad *string, Soyad *string) {
    if Ad == nil {
        panic("Ad nil olamaz")
    }
    if Soyad == nil {
        panic("Soyad nil olamaz")
    }
    fmt.Printf("%s %s\n", *Ad, *Soyad)
    fmt.Println("TamIsim fonksiyonu bitti")
}

func main() {
    Ad := "Yusuf"
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Panik Yok : ", r)
        }
    }()
    TamIsim(&Ad, nil)
    fmt.Println("Ana fonksiyon da bitti")
}

//Çıktımız burada :
//Panik Yok : Soyad nil olamaz

*/