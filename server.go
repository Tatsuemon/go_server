package main

import (
	"fmt"
	"net/http"

	"github.com/Tatsuemon/go_server/models"
	"github.com/codegangsta/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		user := models.NewUser("Naoyoshi Aikawa", 29)
		fmt.Fprintln(w, user)
	})
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
