package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	// "strings"
)

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// fmt.Println(r.Form)
	// fmt.Println(w, "PATH:", r.URL.Path)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadGateway)
	fmt.Fprintf(w, "hello world")
}
func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	switch r.Method {
	case "GET":
		token := time.Now().Unix()
		fmt.Println(token)
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	case "POST":
		if len(r.Form["username"][0]) == 0 {
			fmt.Fprintln(w, "username is empty")
		}
		fmt.Println("==============post=================")
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
