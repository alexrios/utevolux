package simple

const Main = `package main

import (
	"{{.Module}}/api"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	api.Routes(mux)
	log.Fatal(http.ListenAndServe(":1337", mux))
}
`
