// package main

// import (
//     "fmt"
//     "net/http"
// )

// func main() {
//     http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//         http.ServeFile(w, r, "login.html")
//     })

//     fmt.Println("Server started at http://localhost:8080")
//     if err := http.ListenAndServe(":8080", nil); err != nil {
//         fmt.Println("Error starting server:", err)
//     }
// }

package main

import (
	_ "database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// var DB *sql.DB

func main() {

	err := OpenDatabase()
	if err != nil {
		log.Printf("A new error %v", err)
	}
	defer CloseDatabase()
	fmt.Println("Done...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			postLogin(w, r)
		} else {
			http.ServeFile(w, r, "login.html")
		}
	})

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			postSignup(w, r)
		} else {
			http.ServeFile(w, r, "signup.html")
		}
	})

	fmt.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Fprintf(w, "Received: Username=%s, Password=%s", username, password)
}

func postSignup(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	fmt.Fprintf(w, "Received: Username=%s, Password=%s , Email=%s", username, password, email)
}

// func OpenDatabase() error {
// 	var err error
// 	DB, err = sql.Open("postgres", "user=postgres password=8976 dbname=FirstDemoDatabase sslmode=disable")
// 	if err != nil {
// 		return err
// 	}

// 	if err != nil {
// 		return err
// 	}

// 	// CreateTable()
// 	return nil

// }

// func CloseDatabase() error {
// 	return DB.Close()
// }
