package main

import (
	"net/http"
	"html/template"
	"io"
	"github.com/labstack/echo"
)

type M map[string]interface{}

type Renderer struct {
	template *template.Template
	debug bool
	location string
}

/*
	METHOD PARSING TEMPLATE
 */
func (t *Renderer) ReloadTemplates(){
	t.template = template.Must(template.ParseGlob(t.location))
}

/*
	FUNC RENDER TEMPLATE
*/
func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplates()
	return tpl
}

/*
	METHOD EXECUTE TEMPLATE
 */
func (t *Renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}

func main(){
	e := echo.New()

	e.Renderer = NewRenderer("./*.html", true)

	e.GET("/index", func(c echo.Context) error {
		data := M{"message": "Hello World!"}
		return c.Render(http.StatusOK, "index.html", data)
	})

	e.Logger.Fatal(e.Start(":9000"))
}

