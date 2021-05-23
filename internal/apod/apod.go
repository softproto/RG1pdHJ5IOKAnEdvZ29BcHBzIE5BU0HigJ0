package apod

import "encoding/json"

const URL = "https://api.nasa.gov/planetary/apod"

type apodResponse struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	Url            string `json:"url"`
}

type apodErrorResponse struct {
	Code           string `json:"code"`
	Message        string `json:"msg"`
	ServiceVersion string `json:"service_version"`
}

func UnmarshallResponse(response []byte) (resp *apodResponse, err error) {
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func UnmarshallErrorResponse(response []byte) (resp *apodErrorResponse, err error) {
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
