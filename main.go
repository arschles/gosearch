package main

import (
	"log"
	"net/http"
	"os"

	"github.com/foolin/goview"
	echoview "github.com/foolin/goview/supports/echoview-v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// for next time
// - make a UI for all this stuff!
// - caching the results for a few mins or something?
// - maybe proactively getting search results in the background
// - doing image, video, etc... search?
// - deploy this to the actual internet
//		- rate limiting will be needed
// - WRITE OUR OWN CRAWLER!!!

func main() {

	token := os.Getenv("BING_SEARCH_KEY")
	if token == "" {
		log.Fatal("BING_SEARCH_KEY not found")
	}

	disableCache := false
	prod := os.Getenv("PROD")
	if prod == "true" {
		disableCache = true
	}

	// Echo instance
	e := echo.New()
	e.Renderer = echoview.New(goview.Config{
		Root:      "views",
		Extension: ".html",
		Master:    "layouts/base",
		Funcs:     make(map[string]interface{}),
		// TODO: set DisableCache to false when in production
		DisableCache: disableCache,
		Delims: goview.Delims{
			Left:  "{{",
			Right: "}}",
		},
	})

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// curl localhost:12334/api/search?term="how+do+you+mine+bitcoin"

	// TODOs:
	//
	// - figure out what's going on with ACI deploy!
	// - get a favicon (https://i.imgur.com/YbHs016.png courtesy rockerBO)
	// - Need to "fingerprint" the JS and CSS files to
	// 	ensure the browser doesn't cache them after I make
	// 	a change
	// - better styling
	// - little spinner bar after you submit search, before
	//	results are ready
	// - autocomplete & typeahead suggestions!
	// - duckduckgo search operators and "bangs"
	//		- https://help.duckduckgo.com/duckduckgo-help-pages/results/syntax/
	//		- https://duckduckgo.com/bang
	// - maps & location support
	// - duckduckgo-style (google/bing/etc... does it too) "smart sidebar"
	//	tries to guess something quick that you want & puts it just
	// 	to the right of the search results
	// - build the image in ACR tasks and host it in ACR
	// - Get a better DNS for the hostname (we have the ACI hostname now, and we have the front door gosearching.azurefd.net one)
	e.Static("/static", "frontend/public/build")

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", echo.Map{
			"title": "Index title!",
			"add": func(a int, b int) int {
				return a + b
			},
		})
	})
	g := e.Group("/api")
	g.GET("/search", newSearchHandler(token))

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
