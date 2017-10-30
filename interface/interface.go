package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type IPAddr [4]byte

func (ipaddr IPAddr) Sting() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])
}

func main() {
	var a Abser
    f := MyFloat(-math.Sqrt(2))
    v := Vertex{3, 4}

    a = f
    a = &v

    fmt.Println(a.Abs())

    addrs := map[string]IPAddr {
    	"loopback": {127, 0, 0, 1},
    	"googleDNS": {8, 8, 8, 8},
	}

	for n, a:=range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}

type MyFloat float64
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}