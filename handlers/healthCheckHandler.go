package handlers

import (
	"fmt"
	"net/http"
)

func healthCheckHandler(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	fmt.Fprintf(writer, "Server is healthy!")
}
