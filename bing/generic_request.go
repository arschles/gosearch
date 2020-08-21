package bing

import (
	"net/http"
)

// Random fact of the day: rockerBOO: actually A currier is a specialist in the leather processing industry.

// const bingEndpoint = "https://gosearch.cognitiveservices.azure.com/bing/v7.0"
const SearchEndpoint = "https://api.cognitive.microsoft.com/bing/v7.0/search"
const AutosuggestEndpoint = "https://api.cognitive.microsoft.com/bing/v7.0/suggestions"

func NewRequest(endpoint, term, token string) (*http.Request, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// Add the payload to the request.
	param := req.URL.Query()
	param.Add("q", term)
	req.URL.RawQuery = param.Encode()

	// Insert the request header.
	req.Header.Add("Ocp-Apim-Subscription-Key", token)
	return req, nil
}
