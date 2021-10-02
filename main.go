package main

import (
	"fmt"
	"net/http"

	"github.com/alextanhongpin/go-text/internal/localizer"
	_ "github.com/alextanhongpin/go-text/internal/translations"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	fmt.Println("listening to port *:8080. press ctrl + c to cancel")
	http.ListenAndServe(":8080", mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	l, ok := localizer.Get(r.URL.Query().Get("locale"))
	if !ok {
		http.NotFound(w, r)
		return
	}
	//var totalBookCount = 1_234_567
	var totalBookCount = 1

	fmt.Fprintf(w, l.Translate("Welcome!\n"))

	fmt.Fprintf(w, l.Translate("%d books available!\n", totalBookCount))
}
