package httplib

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/xiongdashan/galaxylib"
	"github.com/xiongdashan/galaxylib/dblib"
)

type GApp struct {
	App *echo.Echo
}

func NewGApp() *GApp {
	app := echo.New()
	app.Use(middleware.CORS())
	return &GApp{app}
}

func (g *GApp) UseGContext() {
	// use gcontext by default dbfactory
	g.App.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			gc := NewGContext(c, dblib.DefaultDbFactory())
			return h(gc)
		}
	})

	g.App.Use(middleware.BodyDump(func(c echo.Context, rqbody, rsbody []byte) {
		gc := c.(*GContext)
		gc.Db.Close()
	}))
}

func (g *GApp) AddRoute(path string, fn func(gc *GContext) error) {
	g.App.POST(path, func(ctx echo.Context) error {
		gc := ctx.(*GContext)
		return fn(gc)
	})
}

func (g *GApp) Start() {
	g.App.Logger.Fatal(g.App.Start(galaxylib.GalaxyCfgFile.MustValue("app", "port")))
}
