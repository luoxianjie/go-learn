package main

import "fmt"

const MAX int = 3;

func main(){

    numbers := [MAX]int{1,2,5}
    var i int
    var p [MAX]*int

    for i = 0; i < MAX; i++ {
        p[i] = &numbers[i]
        fmt.Println(*p[i])
    }
}