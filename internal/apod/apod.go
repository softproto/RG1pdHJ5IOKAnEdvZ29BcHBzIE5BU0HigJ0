package apod

import "encoding/json"

//Apod server Response
type Response struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	HdURL          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
}

func UnmarshallResponse(response []byte) (resp *Response, err error) {
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
