package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arschles/gosearch/bing"
	"github.com/labstack/echo/v4"
)

//  curl "localhost:8080/api/search?term=thing"

func newSearchHandler(token string) echo.HandlerFunc {
	type result struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	type queryResults struct {
		ResultList []result `json:"results"`
	}

	appendResult := func(results *queryResults, name, url string) {
		results.ResultList = append(results.ResultList, result{
			Name: name,
			URL:  url,
		})
	}

	return func(ctx echo.Context) error {
		term := ctx.QueryParam("term")
		if term == "" {
			return newError(
				ctx,
				http.StatusBadRequest,
				"No search term in query string",
			)
		}
		// start making the request to Bing
		// Declare a new GET request.
		req, err := bing.NewRequest(bing.SearchEndpoint, term, token)

		// Send the request to Bing.
		resp, err := bing.Client.Do(req)
		if err != nil {
			return newError(
				ctx,
				http.StatusInternalServerError,
				"%s",
				err,
			)
		}

		// Close the response.
		defer resp.Body.Close()

		// Create a new answer.
		ans := &BingAnswer{} // same thing as new(BingAnswer)

		// TODO: this code is from the example site:
		// (https://docs.microsoft.com/en-us/azure/cognitive-services/bing-web-search/quickstarts/go#handle-the-response)
		// it's wrong because passing in &ans to json.Unmarshal
		// passes in a pointer to a pointer - we should just be
		// passing in a pointer.
		// err = json.Unmarshal(body, &ans)

		if err := json.NewDecoder(resp.Body).Decode(ans); err != nil {
			return newError(
				ctx,
				http.StatusInternalServerError,
				"%s",
				err,
			)
		}

		results := &queryResults{}
		// Iterate over search results and print the
		// result name and URL.
		for _, result := range ans.WebPages.Value {
			if result.IsFamilyFriendly {
				appendResult(
					results,
					result.Name,
					result.URL,
				)
			}
		}

		// 	ResultList: []string{
		// 		"definatelyevil",
		// 		"Joker_Dan",
		// 		"bobbingbaboon",
		// 		"rockerBOO",
		// 		"isiahvander",
		// 		"erikdotdev",
		// 		"zanuss",
		// 	},
		// }

		return ctx.JSONPretty(http.StatusOK, results, "  ")
	}
}

func newAutocompleteHandler(token string) echo.HandlerFunc {

	return func(ctx echo.Context) error {
		term := ctx.QueryParam("term")
		if term == "" {
			return newError(
				ctx,
				http.StatusBadRequest,
				"No search term in query string",
			)
		}
		req, err := bing.NewRequest(bing.AutosuggestEndpoint, term, token)
		if err != nil {
			return err
		}

		req.URL.Query().Add("q", term)

		// Send the request to Bing.
		resp, err := bing.Client.Do(req)
		if err != nil {
			return newError(
				ctx,
				http.StatusInternalServerError,
				"%s",
				err,
			)
		}

		// Close the response.
		defer resp.Body.Close()

		autocompleteSuggestions := new(bing.Suggestions)
		if err := json.NewDecoder(resp.Body).Decode(autocompleteSuggestions); err != nil {
			return newError(
				ctx,
				http.StatusInternalServerError,
				"%s",
				err,
			)
		}

		log.Printf("autocomplete suggestions: %+v", *autocompleteSuggestions)
		webSuggestionGroup := new(bing.SuggestionGroup)
		for _, suggestionGroup := range autocompleteSuggestions.SuggestionGroups {
			if suggestionGroup.Name == "Web" {
				*webSuggestionGroup = suggestionGroup
				break
			}
		}
		if webSuggestionGroup == nil {
			return newError(
				ctx,
				http.StatusNotFound,
				"no thing found. red alert!!",
			)
		}

		return ctx.JSON(http.StatusOK, &webSuggestionGroup)
	}
}
