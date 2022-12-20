package main

import "fmt"
//bu bir anonim struct tanımıdır. 
func main(){
	kişi:=struct {
		ad, soyad string
	}{"Melike","Keçelioğlu"}
	fmt.Println(kişi)
}

//bu da bildiğimiz sıradan  bir struct
/*

type insan struct {
	ad, soyad string
}
func main(){
	kişi:= insan{"Melike","Keçelioğlu"}	
	fmt.Println(kişi)
}

*/