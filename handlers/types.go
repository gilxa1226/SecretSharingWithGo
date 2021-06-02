package handlers

type createSecretPayload struct {
	PlainText string `json:"plain_text"`
}

type createSecretResponse struct {
	Id string `json:"id"`
}

type getSecretResponse struct {
	Secret string `json:"data"`
}

type secretData struct {
	Id        string
	PlainText string
}
