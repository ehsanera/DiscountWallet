package models

type ApiResponse struct {
	Bytes      []byte
	StatusCode int
}

type PhoneNumbersResponse struct {
	PhoneNumbers []string `json:"phoneNumbers"`
}
