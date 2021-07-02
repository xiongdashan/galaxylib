package httplib

import "github.com/xiongdashan/galaxylib"

// IHttpConfig 配置接口...
type IHttpConfig interface {
	JwtSigningKey() string

	ListenPort() string
}

// DefaultHTTPCfg for default...
type DefaultHTTPCfg struct {
}

// NewHTTPCfg instance..
func NewHTTPCfg() *DefaultHTTPCfg {
	return &DefaultHTTPCfg{}
}

// JwtSigningKey for jwt sign..
func (d *DefaultHTTPCfg) JwtSigningKey() string {
	return galaxylib.GalaxyCfgFile.MustValue("app", "signingKey")
}

// ListenPort for app port
func (d *DefaultHTTPCfg) ListenPort() string {
	return galaxylib.GalaxyCfgFile.MustValue("app", "port")
}
