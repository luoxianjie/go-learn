package main

import "fmt"

func main(){
    var a int = 10
    var b int = 20
    var c int

    c = max(a,b)

    fmt.Println(c)
}

func max(a,b int) int{
    if a > b{
        return a
    }else{
        return b
    }
}

