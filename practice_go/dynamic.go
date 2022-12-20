package main

import "fmt"

//dinamik atama yapğabilmek için boş. bir inteface e ihtiyacımız var

type dynamic interface{
    
}

func main(){
    //x değişkenimizn tipini interface olarak belirledik
    var x dynamic
    
    x=13
    
    fmt.Println("%T:%v\n", x, x)
    //x değişkenimizin tip ve değerini yazdılarım
    
    x= "selam"
    
     fmt.Println("%T:%v\n", x, x)
}