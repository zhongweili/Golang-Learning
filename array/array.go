package main

import "fmt"

func main() {
    var a [2]string
    a[0] = "hi"
    a[1] = "hao"
    fmt.Println(a)
    primes := [6]int{2, 3, 5 ,7 ,11, 13}
    var s []int = primes[1:4]
    fmt.Println(s)
    
    names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	c := names[0:2]
	d := names[1:3]
	fmt.Println(c, d)

	d[0] = "XXX"
	fmt.Println(c, d)
	fmt.Println(names)

    q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	t := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(t)
}
