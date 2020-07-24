package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func newError(
	ctx echo.Context,
	code int,
	fmtString string,
	params ...interface{},
) error {
	errString := fmt.Sprintf(fmtString, params...)
	ctx.Logger().Print(errString)
	return ctx.JSON(
		code,
		map[string]string{"error": errString},
	)
}
