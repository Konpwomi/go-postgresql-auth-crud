package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		fmt.Fprintln(w, "GET request received")
// 	} else if r.Method == http.MethodPost {
// 		fmt.Fprintln(w, "POST request received")
// 	} else {
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}
// }

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "Hello, JSON!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "About Page")
// }

func main() {
	http.HandleFunc("/json", jsonHandler)
	fmt.Println("Server run at port : 8080")
	http.ListenAndServe(":8080", nil)
} 