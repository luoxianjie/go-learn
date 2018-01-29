//包申明
package main

//导入包
import "fmt"

func main() {
    /**
    * 1, 申明的变量要使用
    * 2, := 赋值的变量之前未申明
    * 3, var,type 都不是强制的
    */
    var a int = 10
    const b = 20
    c := 30

    fmt.Println("a",a,b,c);
}