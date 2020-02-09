package main

import (
	"net/http"

	"github.com/memochou1993/thesaurus/router"
)

func main() {
	r := router.NewRouter()

	http.ListenAndServe(":8081", r)
}
