// package api

// import (
// 	"encoding/json"
// 	"net/http"
// )

// type Item struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }

// var items = []Item{
// 	{ID: 1, Name: "Item 1"},
// 	{ID: 2, Name: "Item 2"},
// 	{ID: 3, Name: "Item 3"},
// }

//	func ReadItemsHandler(w http.ResponseWriter, r *http.Request) {
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(items)
//	}
package api

import (
	"fmt"
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact us at contact@example.com")
}
