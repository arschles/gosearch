package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Handler
func hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello, World!")
}

type searchQuery struct {
	// This has to be exported for the underlying echo query
	// string decoder
	Term string `query:"term"`
}

type queryResults struct {
	ResultList []string
}

// localhost:123/api/search?term=thing

func search(ctx echo.Context) error {
	q := &searchQuery{}
	if err := ctx.Bind(q); err != nil {
		return fmt.Errorf("Couldn't decode query string (%s)", err)
	}

	results := &queryResults{
		ResultList: []string{
			"definatelyevil",
			"Joker_Dan",
			"bobbingbaboon",
			"rockerBOO",
			"isiahvander",
			"erikdotdev",
			"zanuss",
		},
	}

	return ctx.JSONPretty(http.StatusOK, results, "  ")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// curl localhost:12334/api/search?term="how+do+you+mine+bitcoin"

	g := e.Group("/api")
	g.GET("/search", search)

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
