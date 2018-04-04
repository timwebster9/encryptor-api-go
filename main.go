package main

import (
    "fmt"
    "html/template"
    "net/http"
	"log"
	"github.com/timwebster9/encryptor"
)

type CipherData struct {
	Password   string
	Secret     string
	Ciphertext string
}

func start(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
		execMainPage(w)
    }
}

func postAPI(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "POST" {
		var secret = r.FormValue("password")
		var key = r.FormValue("secret")

		encrypted, err := encryptor.Encrypt(secret, key)
		if (err != nil) {
			panic(err)
		}

		cipherData := CipherData{secret, key, encrypted}
		fmt.Printf("Encrypted secret is: %x", cipherData.Ciphertext)

		t, _ := template.ParseFiles("static/index.html")
		t.Execute(w, cipherData)
	}
}

func execMainPage(w http.ResponseWriter) {
	t, _ := template.ParseFiles("static/index.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", start) // set router
	http.HandleFunc("/api", postAPI) // set router

    err := http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}