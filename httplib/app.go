package httplib

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/xiongdashan/galaxylib"
)

// GApp Galaxy Appliction
type GApp struct {
	App *echo.Echo
}

// NewGApp instance...
func NewGApp() *GApp {
	app := echo.New()
	app.Use(middleware.CORS())
	return &GApp{app}
}

// UseGContext 使用带有数据库
func (g *GApp) UseGContext(gc echo.Context) {
	// use gcontext by default dbfactory
	g.App.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return h(gc)
		}
	})

	g.App.Use(middleware.BodyDump(func(c echo.Context, rqbody, rsbody []byte) {
		gc := c.(*GContext)
		gc.Db.Close()
	}))
}

// AddRoute 路由
func (g *GApp) AddRoute(path string, fn func(gc *GContext) error) {
	g.App.POST(path, func(ctx echo.Context) error {
		gc := ctx.(*GContext)
		return fn(gc)
	})
}

// Start 启动...
func (g *GApp) Start() {
	g.App.Logger.Fatal(g.App.Start(galaxylib.GalaxyCfgFile.MustValue("app", "port")))
}
