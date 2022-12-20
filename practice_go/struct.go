package main

import "fmt"

type kişi struct {
    isim string
    soyİsim string
    yaş int
}



func main() {

    kişi1 :=kişi{"Melike", "Keçelioğlu", 21}
    
    fmt.Println(kişi1)
    
    kişi1.isim = "MelikeİsimDeğiştirdi"
    kişi1.soyİsim =  "MelikeEvlendi"
    kişi1.yaş = 12345
    fmt.Println(kişi1) //melike artık çok farklı biri
    
    
    // eğer ki nesne örneği oluşturuyorsak parametreleri boş bırakıp sonradan da atama işlemini gerçekleştirebiliriz
    
    kişi2 := kişi{}
    kişi2.isim, kişi2.soyİsim = "MelikeArtıkbirNesneÖrneğiÜzerindenAtamaYaptırıldı", "İştebudaonunsoyadı"
    kişi2.yaş =98765
    fmt.Println(kişi2)
    
    //veya nesneye özel değişkenleri tanımlarken değiken ismini belirterek de tanımlanama yapabiliriz.
    
   // kişi1:=kişi{soyİsim:"Keçeli", isim:"melike", yaş:21}
    //fmt.Println(kişi1)
    
    
    
    
    
    
    
    
    
    
    
    
    
    

}
















