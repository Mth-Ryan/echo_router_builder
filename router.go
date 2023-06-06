package router_builder

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

type Renderer struct {
	templates *template.Template
}

func getPathsRec(baseFolder, ext string) []string {
	paths := []string{}

	err := filepath.Walk(baseFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, ext) {
				paths = append(paths, path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	return paths
}

func NewRenderer(baseFolder, ext string) *Renderer {
	paths := getPathsRec(baseFolder, ext)
	templates := template.Must(template.ParseFiles(paths...))
	return &Renderer{templates}
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type RouterBuilder struct {
	inner *echo.Echo
}

func NewBuilder(middlewares ...echo.MiddlewareFunc) *RouterBuilder {
	inner := echo.New()
	inner.Use(middlewares...)

	return &RouterBuilder{inner}
}

func parseRoute(r string) string {
	if r == "/" || r == "" {
		return ""
	} else if !strings.HasPrefix(r, "/") {
		return "/" + r
	}
	return r
}

func (b *RouterBuilder) Register(c *Controller) *RouterBuilder {
	g := b.inner.Group(parseRoute(c.baseRoute), c.middlewares...)

	for _, h := range c.handlers {
		g.Add(h.method, parseRoute(h.route), h.inner, h.middlewares...)
	}

	return b
}

func (b *RouterBuilder) RegisterStatic(route, folder string) *RouterBuilder {
	b.inner.Static(parseRoute(route), folder)
	return b
}

func (b *RouterBuilder) RegisterViews(baseFolder, ext string) *RouterBuilder {
	b.inner.Renderer = NewRenderer(baseFolder, ext)
	return b
}

func (b *RouterBuilder) NotFoundHandler(h echo.HandlerFunc) *RouterBuilder {
	b.inner.RouteNotFound("/*", h)
	return b
}

func (b *RouterBuilder) NotFoundView(view string, data any) *RouterBuilder {
	b.NotFoundHandler(func(c echo.Context) error {
		return c.Render(http.StatusOK, view, data)
	})
	return b
}

func (b *RouterBuilder) Build() *echo.Echo {
	return b.inner
}
