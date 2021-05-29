package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerSetup(t *testing.T) {
	mux := http.NewServeMux()
	//SetupHandlers(mux)
	mux.HandleFunc("/healthCheck", healthCheckHandler)
	mux.HandleFunc("/", secretHandler)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/healthCheck", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	if writer.Body.String() != "Server is healthy!" {
		t.Error("Server did not respond as expected")
	}
}
