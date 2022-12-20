package main

import (
    "fmt"
    "time"
    )

func main() {
   k:= make(chan bool, 2) 
   //burada iki adet değer taşıyan bir adet chanel oluşturuyoruz,
   
   // burada time.Sleep ile biri her 5de bir diğeri de her 2de bir ilerleyecek olan iş parçacıklaırnı görebilmek mümkün.
  go func(){
            //burada 5 saniye bekletiyoruz
      time.Sleep(time.Second * 5)
      //burada da k kanalına boolean bir değer gönderiyoruz
      k <- true
      //tekrar 2 sn beklemeye alıyoruz
      
      time.Sleep(time.Second * 2)
      //ve k kanalına tekrar 2.değer görnriliyor
      k <- false
  } ()
  
  //ana iş parçacığımız burada sona ererken k kanalına 2 değer gelene kadar bekleyecek ve işte çıktı
  
  fmt.Println(<-k, <-k)
//burada çıktımız da true ve false olmalı  

   
}