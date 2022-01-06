package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/go-playground/validator/v10"
	"net/http"
)

/*
	STRUCT USER
 */
type User struct {
	Name string `json:"name" form:"name" query:"name" validate:"required"`
	Email string `json:"email" form:"email" query:"email" validate:"required,email"`
	Age int `json:"age" form:"age" query:"age" validate="gte=0,lte=80"`
}

/*
	VALIDATE DATA PAYLOAD REQUEST
 */
type CustomValidator struct {
	validator *validator.Validate
}

/*
	SET VALIDATE TO STRUCT
 */
func (cv *CustomValidator) Validate (i interface{}) error {
	return cv.validator.Struct(i)
}



func main(){
	r := echo.New()

	/*
		PROCESS VALIDATE DATA REQUEST
	 */
	r.Validator = &CustomValidator{validator: validator.New()}

	/*
		TESTING REQUEST PAYLOAD WITH ANY PAYLOAD REQUEST
 	*/
	r.Any("/user", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}

		return c.JSON(http.StatusOK, u)
	})

	/*
		TESTING REQUEST PAYLOAD WITH ANY PAYLOAD REQUEST
	*/
	r.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		if err := c.Validate(u); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, true)
	})



	fmt.Println("server started at 9000")
	r.Logger.Fatal(r.Start(":9000"))
}

/*
	REQUEST CANT PAYLOAD JSON , FORM DATA , XML
 */