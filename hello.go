package main

import (
    "fmt"
    "reflect"
)

func main() {
    var a int
    var b  = '0'
    c := 20
    fmt.Printf("a is a type of %v , b is type of %v , c is type of %v\n", 
        reflect.TypeOf(a) ,reflect.TypeOf(b) ,reflect.TypeOf(c));
    test()
}

func test(){
    fmt.Println("hello world");
}
