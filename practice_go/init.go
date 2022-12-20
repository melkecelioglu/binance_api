package main

import "fmt"

//normalde program çalıştığında ilk olarak main fonksiyonu çalışır
//ama init fonksiyonu varsa önce init fonksiyonu çalışır
//init fonksiyonu bir paket içinde birden fazla olabilir
//init fonksiyonu bir paket içindeki tüm dosyalarda çalışır
//görelim. Bu örnekte global tanımlanmış değişkenin değerini init() fonksiyonunda değiştirdiğimizde main() gibi farklı fonksiyonlarda kullanabildiğimizi göreceğiz.
var degisken string
func init(){
	degisken= "Merhaba"
    fmt.Println("init fonksiyonu yüklendi")
    
}

func main(){
    	fmt.Println(degisken)
		fmt.Println("main fonksiyonu yüklendi")
	
	//pointerli kisim
		degistirr(&degisken)

	}

// benzer islemi pointerlarla da yaptirabilmemiz mumkun

func degistirr(degisken *string){
	*degisken= "merhaba"
}



