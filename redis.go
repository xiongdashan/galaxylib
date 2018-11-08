package galaxylib

import (
	"encoding/json"

	"github.com/go-redis/redis"
)

type GalaxyRedisClient struct {
}

var DefaultGalaxyRedisClient = &GalaxyRedisClient{}

func (self GalaxyRedisClient) Get(key string, v interface{}) (err *GalaxyError) {
	//cfg := GetDefaultConfig()
	host := GalaxyCfgFile.MustValue("redis", "host")
	prefix := GalaxyCfgFile.MustValue("redis", "prefix")
	client := redis.NewClient(&redis.Options{
		Addr: host,
	})
	//key = fmt.Sprintf("%s")
	result := client.Get(prefix + key)

	//return client
	if e := result.Err(); e != nil {
		err = DefaultGalaxyError.FromError(1, e)
		return
	}

	buffer, e := result.Bytes()

	if e != nil {
		err = DefaultGalaxyError.FromError(1, e)
		return
	}

	if len(buffer) == 0 {
		err = DefaultGalaxyError.FromText(1, "empty")
	}

	if e := json.Unmarshal(buffer, v); e != nil {

		err = DefaultGalaxyError.FromError(1, e)
		return
	}
	return
}
