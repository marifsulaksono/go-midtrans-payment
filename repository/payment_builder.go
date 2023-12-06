package repository

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func RequestMidtransHitter(key, link string, payload any) (*http.Response, error) {
	payloadRequest, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, link, bytes.NewBuffer(payloadRequest))
	if err != nil {
		return nil, err
	}

	// header set-up
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic "+key)

	// hit midtrans API enpoint with the prepared request
	client := http.Client{}
	response, err := client.Do(request)

	return response, err
}
