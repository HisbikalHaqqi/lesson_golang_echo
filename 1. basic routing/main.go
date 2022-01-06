package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

type M map[string]interface{}

func main() {
	r := echo.New()

	/* ======================================== ROUTING ECHO FRAMEWORK =================================== */
	/*
		METHOD CTX.STRING = PLAIN TEXT
	 */
	r.GET("/", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})

	/*
		METHOD CTX.HTML = RENDER HTML
	 */
	r.GET("/html", func(ctx echo.Context) error{
		data := "Hello from html"
		return ctx.HTML(http.StatusOK,data)
	})

	/*
		METHOD CTX.REDIRECT = REDIRECT ROUTE
	 */
	r.GET("/redirect", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	})

	/*
		METHOD CTX.JSON = RENDER JSON
	 */
	r.GET("/json", func(ctx echo.Context) error {
		data := M{"Message":"Hello","Counter":2}
		return ctx.JSON(http.StatusOK, data)
	})

	/*
		METHOD QUERY PARAM = GET PARAM
	 */
	r.GET("/parsing-string", func(ctx echo.Context) error{
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("Hello %s", name)

		return ctx.String(http.StatusOK, data)
	})

	/*
		METHOD PARAM = PATH PARAMETER SCHEMA ROUTE (URL SEGMENT 2)
	 */
	r.GET("/param/:name", func(ctx echo.Context) error{
		//segment 2
		name := ctx.Param("name")
		data := fmt.Sprintf("Hello %s", name)
		return ctx.String(http.StatusOK, data)
	})

	/*
		METHOD URL PATH PARAM SETELAHNYA (URL SEGMENT 3 DST)
	 */
	r.GET("/page3/:name/*", func(ctx echo.Context) error {
		name := ctx.Param("name")
		message := ctx.Param("*")

		data := fmt.Sprintf("Hello %s, I have message for you: %s", name, message)

		return ctx.String(http.StatusOK, data)
	})

	/*
		PARSING FORM DATA
	 */
	r.POST("/page4", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")

		data := fmt.Sprintf(
			"Hello %s, I have message for you: %s",
			name,
			strings.Replace(message, "/", "", 1),
		)

		return ctx.String(http.StatusOK, data)
	})

	/* ======================================== ROUTING ECHO FRAMEWORK =================================== */

	/* ======================================== ROUTING WITH WRAP HANDLER =================================== */
	var ActionIndex = func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("from action index"))
	}

	var ActionHome = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("from action home"))
	})

	r.GET("/index", echo.WrapHandler(http.HandlerFunc(ActionIndex)))
	r.GET("/home", echo.WrapHandler(ActionHome))
	/* ======================================== ROUTING WITH WRAP HANDLER =================================== */

	/* ======================================== ROUTING STATIC =================================== */
	r.Static("/static","assets")
	/* ======================================== ROUTING STATIC =================================== */



	r.Start(":9000")
}