package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Person struct {
	Address Address
}

type Address struct {
	DoorNbr    int
	StreetName string
	City       string
}

func (address Address) tostring() string {
	return strconv.Itoa(+address.DoorNbr) + " " + address.StreetName + " " + address.City
}
func showAddress(p Person) {
	fmt.Println(p.Address.tostring())
}

func main() {
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(rw, "Hello %q!", html.EscapeString(r.URL.Path))
	// })

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		http.ServeFile(rw, r, "htmls"+r.URL.Path[1:])
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
