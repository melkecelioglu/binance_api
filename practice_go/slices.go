package main

import "fmt"

func main() {
   a:= [6]int{2,3,4,5,6,7}
   fmt.Println(a)
   var b []int = a[2:4]
   fmt.Println(b)

   var c []int =a[3:]
   fmt.Println(c)
   

   fmt.Println("a uzunluk: ", len(a))
   fmt.Println("a kapasite: ", cap(a))
   fmt.Println("a icerigi: ", a)
   fmt.Println("b uzunluk: ", len(b))
   fmt.Println("b kapasite: ", cap(b))
   fmt.Println("b icerigi: ", b)
   fmt.Println("c uzunluk: ", len(c))

   //make ile dilim oluşturma
   k:= make([]int, 5)
   fmt.Println(k)
   k = append(k,6)
   k = append(k,7)
   fmt.Println(k)
   fmt.Println(len(a), cap(a))
}