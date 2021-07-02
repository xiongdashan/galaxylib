package httplib

import (
	"context"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/xiongdashan/galaxylib/dblib"
)

// JwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Name   string `json:"name"`
	ID     uint   `json:"id"`
	Admin  bool   `json:"admin"`
	Avatar string `json:"avatar"`
	jwt.StandardClaims
}

// IGalaxyApp  application..
type IGalaxyApp interface {
	UseGContext()
	AddRoute(path string, fn func(gc IGContext) error)
	GroupRoute(path string, group *echo.Group, fn func(gc IGContext) error)
	AddGet(path string, fn func(gc IGContext) error)
	NativeRoute(pat string, fn echo.HandlerFunc)
	CreateGroup(name string, jwt bool) *echo.Group
	Start()
	UseJWT()
	Use(fn echo.MiddlewareFunc)
	Stop(ctx context.Context)
}

// GApp Galaxy Appliction
type GApp struct {
	App      *echo.Echo
	GContext IGContext
	cfg      IHttpConfig
	db       dblib.IDbFactory
}

// NewGApp instance...
func NewGApp(app *echo.Echo, db dblib.IDbFactory, cfg IHttpConfig) *GApp {
	app.Use(middleware.CORS())
	app.Use(middleware.Logger())
	app.Debug = true
	return &GApp{
		App: app,
		db:  db,
		cfg: cfg,
	}
}

// Use Middleware..
func (g *GApp) Use(fn echo.MiddlewareFunc) {
	g.App.Use(fn)
}

// UseJWT for authentions..
func (g *GApp) UseJWT() {
	buf, _ := ioutil.ReadFile(g.cfg.JwtSigningKey())
	g.App.Use(middleware.JWT(buf))
}

// CreateGroup for ruoute...
func (g *GApp) CreateGroup(name string, jwt bool) *echo.Group {
	r := g.App.Group(name)

	if jwt {

		// buf, err := ioutil.ReadFile(g.cfg.JwtSigningKey())

		// if err != nil {
		// 	panic(err)
		// }

		config := middleware.JWTConfig{
			Claims:     &JwtCustomClaims{},
			SigningKey: []byte(g.cfg.JwtSigningKey()),
		}
		r.Use(middleware.JWTWithConfig(config))
	}
	return r
}

// // GroupRoute for app
// func (g *GApp) GroupRoute(path string, group *echo.Group, fn func(gc IGContext) error) {
// 	group.POST(path, func(ctx echo.Context) error {
// 		gc := ctx.(IGContext)
// 		return fn(gc)
// 	})
// }

// UseGContext 使用带有数据库
func (g *GApp) UseGContext() {
	// use gcontext by default dbfactory
	g.App.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				g.GContext.CloseDb()
			}()
			g.GContext = NewContext(g.db, c) //.BindContext(c)

			return next(g.GContext)
		}
	})

	// g.App.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		defer func() {
	// 			gc := c.(IGContext)
	// 			gc.CloseDb()
	// 		}()
	// 		return h(c)
	// 	}
	// })
}

// AddRoute 路由
func (g *GApp) AddRoute(path string, fn func(gc IGContext) error) {
	g.App.POST(path, func(ctx echo.Context) error {
		gc := ctx.(IGContext)
		return fn(gc)
	})
}

// AddGet GET 路由
func (g *GApp) AddGet(path string, fn func(gc IGContext) error) {
	g.App.GET(path, func(ctx echo.Context) error {
		gc := ctx.(IGContext)
		return fn(gc)
	})
}

// GroupRoute for group settings..
func (g *GApp) GroupRoute(path string, group *echo.Group, fn func(gc IGContext) error) {

	group.POST(path, func(ctx echo.Context) error {
		gc := ctx.(IGContext)
		return fn(gc)
	})
}

// NativeRoute  to echo
func (g *GApp) NativeRoute(path string, fn echo.HandlerFunc) {
	g.App.GET(path, fn)
}

// Start 启动...
func (g *GApp) Start() {
	g.App.Logger.Info(g.App.Start(g.cfg.ListenPort()))
}

// Stop for app shutdown
func (g *GApp) Stop(ctx context.Context) {
	g.App.Shutdown(ctx)
}
