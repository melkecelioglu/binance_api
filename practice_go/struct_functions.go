package main

import "fmt"

type insan struct{
	isim string
	yas int
}

func(i insan) tanimaFonksiyonu(){
	fmt.Println("Benim adim", i.isim, "yasim", i.yas)
}

func main(){
	kişi:=insan{"Melike", 21}
	kişi.tanimaFonksiyonu()
}