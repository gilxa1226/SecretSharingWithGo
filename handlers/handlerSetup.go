package handlers

import "net/http"

func SetupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/healthCheck", healthCheckHandler)
	mux.HandleFunc("/", secretHandler)
}
