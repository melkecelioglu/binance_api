package main

import "fmt"

func main(){
	var a [3] string
	a[0] = "Melike"
	a[1] = "Keçelioğlu"
	a[2] = "Bu da a[2] aslında"

	fmt.Println(a)
	fmt.Println(a[0])
	b:=[3]string{"Akisisi","Bkisisi","Ckisisi"}

	fmt.Println(b)
}