package main

import "fmt"

    var isimler = []string{"M", "K","H","A"}
func main() {
    for a,b := range isimler { 
        fmt.Printf("%d. indeks = %s\n",a,b)
    }
}