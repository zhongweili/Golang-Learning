package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
	//switch r.URL.Path {
	//case "/string":
	//	fmt.Fprint(w, String("I'm a frayed knot."))
	//case "/struct":
	//	fmt.Fprint(w, &Struct{"Hello", ":", "Gophers!"})
	//}
	fmt.Fprint(w, "world")

}

type String string

type Struct struct {
	Greeting string
	Punct string
	who string
}

func (str String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
		fmt.Fprint(w, str)
}

func (stru Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {
		fmt.Fprint(w, stru.Greeting+stru.Punct+stru.who)
}

func main() {
	var h Hello
	http.Handle("/", h)

//	err := http.ListenAndServe("localhost:4000", h)

	http.Handle("/string", String("I'm a frayed knot."))

	http.Handle("/struct", Struct{"Hi", ":", "Gophers!"})

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}