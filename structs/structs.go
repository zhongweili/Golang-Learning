package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1 ,2}
    v.X = 4
    var p = &Vertex{1, 3}
    fmt.Println(v.X, p)
}
