package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("server has start:", time.Now())

	e := echo.New()

	e.Use(middleW)

	e.GET("/", handler)
	err := e.Start(":8080")
	if err != nil {
		e.Logger.Fatal(err)
	}
}

func handler(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, dayBefore())
	if err != nil {
		return err
	}
	return nil
}

func dayBefore() string {
	timeWeNeed := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	dur := time.Until(timeWeNeed)
	return fmt.Sprintf("Days until to 1 January 2025: %v", int(dur.Hours()/24))

}

func middleW(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")
		if val == "admin" {
			log.Println("red button user detected")
		}
		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
