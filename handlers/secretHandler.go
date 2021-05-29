package handlers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var secretMap = inMemoryMap{Mu: sync.Mutex{}, Store: make(map[string]string)}

func secretPostHandler(writer http.ResponseWriter, request *http.Request) {
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Error reading body", http.StatusInternalServerError)
		return
	}
	secret := createSecretPayload{}
	err = json.Unmarshal(bodyBytes, &secret)
	if err != nil || len(secret.PlainText) == 0 {
		http.Error(writer, "Invalid Request", http.StatusBadRequest)
		return
	}
	digest := md5.Sum([]byte(secret.PlainText))
	response := createSecretResponse{Id: fmt.Sprintf("%x", digest)}

	s := secretData{Id: response.Id, PlainText: secret.PlainText}
	secretMap.Write(s)
	jd, err := json.Marshal(&response)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(jd)
}

func secretGetHandler(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Path[1:]
	response := getSecretResponse{}
	secret := secretMap.Read(id)
	response.Secret = secret

	jd, err := json.Marshal(&response)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("content-type", "application/json")
	if len(response.Secret) == 0 {
		writer.WriteHeader(404)
	}
	writer.Write(jd)

}

func secretHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		secretGetHandler(writer, request)
	} else if request.Method == "POST" {
		secretPostHandler(writer, request)
	}
}
