package main

import "fmt"

func main() {

    var a int = 10
    var b int = 21
    var c int

    c = a + b
    fmt.Printf("a + b = %d\n",c)
    c = a - b
    fmt.Printf("a - b = %d\n",c)
    c = a * b
    fmt.Printf("a * b = %d\n",c)
    c = a / b
    fmt.Printf("a / b = %d\n",c)
    c = b % a;
    fmt.Printf("b %% a = %d",c)

}