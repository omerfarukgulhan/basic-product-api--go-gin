package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()

	err := e.Start("localhost:8080")
	if err != nil {
		return
	}
}
