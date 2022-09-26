package handler

import (
	"net/http"
	"database/sql"

	"github.com/julienschmidt/httprouter"
)

type Params struct {
	DB *sql.DB
}

func Router(params *Params) *httprouter.Router {
	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})

	router.GET("/api/test", apiHandlerTest(params))

	return router
}
