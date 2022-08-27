package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type Page struct {
	Title           string
	Author          string
	Header          string
	PageDescription string
	ContentTitle    string
	Content         string
	URI             string
}

func loadFile(fileName string) (string, error) {
	bytess, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}
	return string(bytess), nil
}

func handler(w http.ResponseWriter, r *http.Request) {

	uri := "www.github.com/ysfaltn02/golang"

	page := Page{
		Title:           "Go Lang Load File With Template",
		Author:          "Yusuf ALTUN",
		Header:          "Welcome to My Blog",
		PageDescription: "Go Lang Load File With Template",
		ContentTitle:    "Go Lang Web API",
		Content:         "Lorem ipsum dolor sit amet consectetur adipisicing elit. Iusto, voluptate.",
		URI:             "https://" + uri}

	tmpl, err := template.ParseFiles("page.html")

	if err != nil {
		fmt.Println("Error Occured")
	}

	tmpl.Execute(w, page)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
