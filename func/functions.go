package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func main() {
    var i, j int = 1, 2
    k := 3
    c, python, java := true, false, "no!"
    fmt.Println(i, j, k, c, python, java)
    fmt.Println(add(42,13))
}
