package galaxylib

import "github.com/Unknwon/goconfig"

var GalaxyCfgFile *goconfig.ConfigFile

type GalaxyConfiger struct {
}

func init() {
	DefaultGalaxyConfig.InitConfig()
}

var DefaultGalaxyConfig = &GalaxyConfiger{}

func (c *GalaxyConfiger) InitConfig() {
	if GalaxyCfgFile != nil {
		return
	}
	GalaxyCfgFile = c.GetDefaultConfig()
}

//GetConfig 获取配置文件,default ./config/config.ini
func (c *GalaxyConfiger) GetConfig(path string) (cfg *goconfig.ConfigFile) {
	if GalaxyCfgFile == nil {
		GalaxyCfgFile, _ = goconfig.LoadConfigFile(path)
	}
	cfg = GalaxyCfgFile
	return
}

//GetDefaultConfig 获取配置文件
func (c *GalaxyConfiger) GetDefaultConfig() *goconfig.ConfigFile {
	return c.GetConfig("../config/config.ini")
}
