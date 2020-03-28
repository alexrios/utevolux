package simple

const DummyHandler = `package api

import (
	domain "{{.Module}}"
	"net/http"
)

type DummyHandler struct {
	service domain.DummyService
}

func (h *DummyHandler) index(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("It Works"))
}
`

const Routes = `package api

import (
	"github.com/urfave/negroni"
	"net/http"
)

func Routes(mux *http.ServeMux) {
	n := negroni.Classic()
	n.UseHandler(mux)

	//Routes
	handler := DummyHandler{}
	mux.HandleFunc("/", handler.index)
}`