package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func apiHandlerTest(params *Params) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		jsonEncoder := json.NewEncoder(w)
		jsonEncoder.SetIndent("", "\t")
		jsonEncoder.Encode(struct{}{})
	}
}
