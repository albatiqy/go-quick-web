package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func MiddleCORS(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		header := w.Header()
		header.Set("Access-Control-Allow-Origin", "*")
		next(w, r, ps)
	}
}
