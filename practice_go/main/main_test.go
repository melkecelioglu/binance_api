package main

import "testing"

func TestMerhaba(t *testing.T){
    if Merhaba("Kaan")!= "Selam Kaan" {
        t.Error("Merhaba fonksiyonunda bir sikinti var")
    }
}

//terminalde go test command line gir

// go test
//İçerisinde bulunduğu projenin tüm test fonksiyonlarını test eder.

// go test -v
//Her test için ayrı bilgi verir.

// go test -timeout 30s
//30 saniye zaman asimi ile test eder.

// go test -run TestMerhaba
//sadece belirli bir fonksiyonu test eder

