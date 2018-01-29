package main

import "fmt"

func main (){

    a := 10

    var b int

    c := [6]int{1,2,5,6}

    for b = 0; b<a; b++{
        fmt.Println(b);
    }

    a = 5;
    for a < b {
        fmt.Println(a);
        a++;
    }

    for key,value := range c {
        fmt.Printf("c 第 %d 位 是 %d\n",key,value);
    }


}
