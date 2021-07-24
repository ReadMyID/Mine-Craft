package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Page struct {
	Title string
	Body  []byte
}
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
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func fetchFile() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		http.ServeFile(rw, r, "htmls/"+r.URL.Path[1:])
	})
}

func main() {
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(rw, "Hello %q!", html.EscapeString(r.URL.Path))
	// })

	fetchFile()
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
