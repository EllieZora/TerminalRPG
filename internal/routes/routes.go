package routes

import (
	"net/http"
)

func NewRouter() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/", indexHandler)

	return router
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to TerminalRPG!"))
}
