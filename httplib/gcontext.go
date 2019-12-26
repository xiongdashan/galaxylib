package httplib

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/xiongdashan/galaxylib"

	"github.com/xiongdashan/galaxylib/dblib"

	"github.com/labstack/echo"
)

type GContext struct {
	echo.Context
	Db dblib.IDbFactory
}

func NewGContext(e echo.Context, db dblib.IDbFactory) *GContext {
	g := &GContext{}
	g.Context = e
	g.Db = db
	return g
}

func (g *GContext) Commit() {
	g.Db.Commit()
}

func (g *GContext) Rollback() {
	g.Db.Rollback()
}

func (g *GContext) Train() {
	g.Db.Tran()
}

func (g *GContext) Err(code int, err error) error {
	return g.JSON(http.StatusOK, echo.Map{
		"code": code,
		"msg":  err.Error(),
	})
}

func (g *GContext) ErrMsg(code int, msg ...string) error {
	return g.returnData(code, msg...)
}

func (g *GContext) OK(data interface{}) error {
	return g.JSON(http.StatusOK, data)
}

func (g *GContext) FilterIP(ip ...string) error {
	realip := g.Context.RealIP()

	if len(ip) == 0 {
		ip = galaxylib.GalaxyCfgFile.MustValueArray("app", "ip", ",")
	}

	for _, i := range ip {
		if strings.TrimSpace(i) == realip {
			return nil
		}
	}
	return g.ErrMsg(504, "非法IP请求:%s", realip)
}

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
