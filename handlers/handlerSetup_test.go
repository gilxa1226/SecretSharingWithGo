package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlerSetup(t *testing.T) {
	mux := http.NewServeMux()
	SetupHandlers(mux)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/healthCheck", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	if writer.Body.String() != "Server is healthy!" {
		t.Error("Server did not respond as expected")
	}

	writer = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 404 {
		t.Errorf("Response code is %v", writer.Code)
	}

	if writer.Body.String() != "{\"data\":\"\"}" {
		t.Error("Server did not respond as expected: ", writer.Body.String())
	}

	writer = httptest.NewRecorder()
	request, _ = http.NewRequest("POST", "/", strings.NewReader(""))
	mux.ServeHTTP(writer, request)

	if writer.Code != 400 {
		t.Errorf("Response code is %v", writer.Code)
	}

}
