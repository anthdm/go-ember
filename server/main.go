package main

import "net/http"

func main() {

	http.Handle("/", http.FileServer(http.Dir("public/dist")))

	http.ListenAndServe(":3000", nil)
}

func renderIndex(w http.ResponseWriter, r *http.Request) {

}
