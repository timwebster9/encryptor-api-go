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

		formatted := fmt.Sprintf("%x", encrypted)
		cipherData := CipherData{secret, key, formatted}
		fmt.Printf("Encrypted secret is: %s", cipherData.Ciphertext)

		t, _ := template.ParseFiles("static/index.html")
		t.Execute(w, cipherData)
	}
}

func execMainPage(w http.ResponseWriter) {
	t, err := template.ParseFiles("static/index.html")
	if (err != nil) {
		log.Fatalf("Error parsing template: %s", err)
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", start) // set router
	http.HandleFunc("/api", postAPI) // set router

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) 

    err := http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}