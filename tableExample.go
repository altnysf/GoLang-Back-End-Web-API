package main

import (
	"fmt"
	"net/http"
)

type Human struct {
	Fname string
	Lname string
	Age   int
}

func (hum Human) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hum.Fname = "TestHumanName"
	hum.Lname = "TestHumanSurname"
	hum.Age = 30

	// To Parse The Form
	r.ParseForm()

	// To Get Form Information From Server
	fmt.Println(r.Form)

	// To Get the Path Information of the URL
	fmt.Println("Path : ", r.URL.Path)

	fmt.Fprint(w, "<style> td{border:1px solid red;border-radius:10px;padding:10px;}</style><table><tr><td><b>Ad</b></td><td><b>Soyad</b></td><td><b>Yaş</b></td></tr><tr><td>", hum.Fname, "</td><td>", hum.Lname, "</td><td>", hum.Age, "</td></tr><tr></tr><tr></tr><tr><td> <b>Path</b></td><td>", r.URL.Path, "</td></tr></table>")

}

func main() {
	var hum Human
	err := http.ListenAndServe("localhost:9000", hum)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
