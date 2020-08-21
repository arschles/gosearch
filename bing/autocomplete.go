package bing

type SearchAction struct {
	DisplayText string `json:"displayText"`
	Query       string `json:"query"`
	SearchKind  string `json:"searchKind"`
	URL         string `json:"url"`
}

type SuggestionGroup struct {
	Name              string         `json:"name"`
	SearchSuggestions []SearchAction `json:"searchSuggestions"`
}
type Suggestions struct {
	SuggestionGroups []SuggestionGroup `json:"suggestionGroups"`
}
