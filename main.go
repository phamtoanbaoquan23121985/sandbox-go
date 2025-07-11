package main

import (
	"fmt"
	"net/http"
	"os"
)

func servePage(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile(filename)
		if err != nil {
			http.Error(w, "Page not found", http.StatusNotFound)
			return
		}
		w.Write(content)
	}
}

func main() {
	http.HandleFunc("/", servePage("home.html"))
	http.HandleFunc("/about", servePage("about.html"))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server running at port 8080")
	http.ListenAndServe(":8080", nil)
}
