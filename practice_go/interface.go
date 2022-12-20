package main

import "fmt"

 type biberSalcasi struct{
     
 }
 
 func(s biberSalcasi) Ye(){
     fmt.Println("Biber Salcasi yenildi")
 }
 
 type domatesSalcasi struct{
     
 }

 type Salca interface{
	 Ye()
	  }
 
 func (s domatesSalcasi) Ye(){
     fmt.Println("Domates salcasi yenildi")
 }
 
 func main(){
     
     biber := biberSalcasi{}
     domates := domatesSalcasi{}
     Ye(biber)
     Ye(domates)
     var salcam Salca
     salcam = &biber
     salcam.Ye()
     
 }
 
 func Ye(s Salca){
     s.Ye()
 }