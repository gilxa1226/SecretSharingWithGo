package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func testGetSecret(t *testing.T) {

	mux := http.NewServeMux()
	SetupHandlers(mux)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/", strings.NewReader(""))
	mux.ServeHTTP(writer, request)

}
