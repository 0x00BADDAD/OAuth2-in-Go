package server

// Aim is to find the way to write an http response to be an html file

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type tmplData struct {
	Title string
	Items []string
}

func myHandler(w http.ResponseWriter, req *http.Request) {
	// writes response
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	cwd, err := os.Getwd()
	fmt.Printf("Value of cwd is: %s\n", cwd)
	if err != nil {
		fmt.Print("Error in os.Getwd()")
	}
	path_to_file := filepath.Join(cwd, "templates", "landing.tmpl")
	tmpl, err := template.ParseFiles(path_to_file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Whoops! something went wrong while parsing the template!")
		fmt.Printf("The error while parsing landing.tmpl is: %v\n", err)
		return
	}
	data := tmplData{
		Title: ">>Title here!<<",
		Items: []string{"Sheldon", "Penny", "Raj", "Howard"},
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Whoops! something went wrong when executing the data on template!")
		fmt.Printf("Err while executing the data on landing.tmpl is %v\n", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func InitServer() {

	http.Handle("/", http.HandlerFunc(myHandler))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("The value of Err when exiting the server: %v\n", err)
	}

}
