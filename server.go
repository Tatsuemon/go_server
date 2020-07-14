package main

import (
	"net/http"

	"github.com/Tatsuemon/go_server/models"
	"github.com/codegangsta/negroni"
	"github.com/unrolled/render"
)

func main() {
	ren := render.New()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		user := models.NewUser("Naoyoshi Aikawa", 29)
		ren.JSON(w, http.StatusOK, user)
	})
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
