package httplib

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/xiongdashan/galaxylib"

	"github.com/xiongdashan/galaxylib/dblib"

	"github.com/labstack/echo/v4"
)

// IGContext 数据上下文
type IGContext interface {
	echo.Context

	// Commit 数据提交
	Commit()

	// Rollback 回滚
	Rollback()

	// Train
	Train()

	// Err
	Err(code int, err error) error

	// ErrMsg
	ErrMsg(code int, msg ...string) error

	// Ok
	OK(data interface{}) error

	// FilterIP
	FilterIP(ip ...string) error

	// FiterKey
	FilterKey(key ...string) error

	// CloseDb
	CloseDb()

	// BindContext
	BindContext(ctx echo.Context)

	DB() dblib.IDbFactory
}

// GContext  Galaxy app 上下文
type GContext struct {
	echo.Context
	db dblib.IDbFactory
}

// NewGContext instance ...
func NewGContext(db dblib.IDbFactory) *GContext {
	g := &GContext{}
	g.db = db
	return g
}

// NewGContext instance ...
func NewContext(db dblib.IDbFactory, ctx echo.Context) *GContext {
	g := &GContext{ctx, db}
	return g
}

// DB lib in context
func (g *GContext) DB() dblib.IDbFactory {
	return g.db
}

// BindContext 绑定
func (g *GContext) BindContext(ctx echo.Context) {
	g.Context = ctx
}

// CloseDb 关闭数据...
func (g *GContext) CloseDb() {
	g.db.Close()
}

// Commit 数据提交
func (g *GContext) Commit() {
	g.db.Commit()
}

// Rollback 数据回滚
func (g *GContext) Rollback() {

	g.db.Rollback()
}

// Train 开户事务
func (g *GContext) Train() {
	g.db.Tran()
}

// Err 返回错误信息
func (g *GContext) Err(code int, err error) error {
	return g.JSON(http.StatusOK, echo.Map{
		"code": code,
		"msg":  err.Error(),
	})
}

// ErrMsg 返回错误信息
func (g *GContext) ErrMsg(code int, msg ...string) error {
	return g.returnData(code, msg...)
}

// OK 正常请求返
func (g *GContext) OK(data interface{}) error {
	return g.JSON(http.StatusOK, data)
}

// FilterIP IP过滤
func (g *GContext) FilterIP(ip ...string) error {

	if len(ip) == 0 {
		ip = galaxylib.GalaxyCfgFile.MustValueArray("app", "ip", ",")
	}

	realip := g.Context.RealIP()

	for _, i := range ip {
		if strings.TrimSpace(i) == realip {
			return nil
		}
	}
	return g.ErrMsg(504, "非法IP请求:%s", realip)
}

// FilterKey 过滤参数..
func (g *GContext) FilterKey(key ...string) error {

	rq := strings.TrimSpace(g.QueryParam("key"))
	//g.Request().URL.RawQuery
	k := ""

	if len(key) == 0 {
		k = galaxylib.GalaxyCfgFile.MustValue("app", "key")
	} else {
		k = key[0]
	}

	if rq == "" || rq != k {
		return g.ErrMsg(505, "非法请求")
	}
	return nil
}

func (g *GContext) returnData(code int, msg ...string) error {
	ret := &struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{
		Code: code,
		Msg:  "",
	}
	if len(msg) == 0 {
		return g.JSON(http.StatusOK, ret)
	}
	if len(msg) == 1 {
		ret.Msg = msg[0]
		return g.JSON(http.StatusOK, ret)
	}
	ret.Msg = fmt.Sprintf(msg[0], msg[:1])
	return g.JSON(http.StatusOK, ret)
}
