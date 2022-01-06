/*
	MIDDLEWARE WITH LOGRUS, ECHO LOGGGER
*/
package main

import (
    "fmt"
    "github.com/labstack/echo"
    "net/http"
)

func main(){
	e := echo.New()

	/*
		PROCESS MIDDLEWARE
	*/
	/* ========================================== MIDDLEWARE WITH SCHEMA ECHO ==================================== */
	e.Use(middlewareOne)
	e.Use(middlewareTwo)
	/* ========================================== MIDDLEWARE WITH SCHEMA ECHO ==================================== */

	/* ========================================== MIDDLEWARE WITH NON SCHEMA ECHO ==================================== */
	e.Use(echo.WrapMiddleware(middlewareSomething))
	/* ========================================== MIDDLEWARE WITH NON SCHEMA ECHO ==================================== */

	e.GET("/",func (c echo.Context) (err error)  {
		fmt.Println("three")

		return c.JSON(http.StatusOK, true)
	})

	e.Logger.Fatal(e.Start(":9000"))
}

/* ========================================== MIDDLEWARE WITH SCHEMA ECHO ==================================== */
/*
	MIDDLEWARE 1
*/
func middlewareOne(next echo.HandlerFunc) echo.HandlerFunc{
	return func(c echo.Context) error {
		fmt.Println("from middleware one")
		return next(c)
	}
}

/*
	MIDDLEWARE 2
*/
func middlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("from middleware two")
		return next(c)
	}
}
/* ========================================== MIDDLEWARE WITH SCHEMA ECHO ==================================== */

/* ========================================== MIDDLEWARE WITH NON SCHEMA ECHO ==================================== */
func middlewareSomething(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("from middleware something")
        next.ServeHTTP(w, r)
    })
}
/* ========================================== MIDDLEWARE WITH NON SCHEMA ECHO ==================================== */