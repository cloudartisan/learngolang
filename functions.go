package main

import "fmt"

func addition(x int, y int) int {
    return x + y
}

func main() {
    var x int = 1
    var y int = 2
    fmt.Printf("%d + %d = %d\n", x, y, addition(x, y))
}
