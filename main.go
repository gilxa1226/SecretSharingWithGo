package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SecretPost struct {
	PlainText string `json:"plain_text"`
}

type PostResponse struct {
	id string `json:"id"`
}

var secretMap = make(map[string]string)

func secretPostHandler(writer http.ResponseWriter, request *http.Request) {
	len := request.ContentLength
	hash := md5.New()

	if len <= 0 {
		writer.WriteHeader(400)
	} else {
		body := make([]byte, len)
		request.Body.Read(body)
		var sp SecretPost
		err := json.Unmarshal(body, &sp)

		if err == nil && sp.PlainText != "" {
			io.WriteString(hash, sp.PlainText)
			strhash := hex.EncodeToString(hash.Sum(nil))
			secretMap[strhash] = sp.PlainText
			writer.Header().Set("content-type", "application/json")
			fmt.Fprintf(writer, "{ \"id\": \"%s\" }", strhash)
		} else {
			writer.WriteHeader(400)
		}
	}
}

func secretGetHandler(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Path[1:]
	secret := secretMap[id]
	writer.Header().Set("content-type", "application/json")
	if len(id) == 0 || secret == "" {
		writer.WriteHeader(404)
		fmt.Fprintf(writer, "{\"data\": \"\"}")
	} else {
		fmt.Fprintf(writer, "{\"data\": \"%s\"}", secret)
	}

}

func secretHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		secretGetHandler(writer, request)
	} else if request.Method == "POST" {
		secretPostHandler(writer, request)
	}
}

func healthCheckHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Server is healthy!")
}

func main() {
	http.HandleFunc("/", secretHandler)
	http.HandleFunc("/healthCheck", healthCheckHandler)
	http.ListenAndServe(":8080", nil)
}
