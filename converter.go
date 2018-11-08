package galaxylib

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

type GalaxyConverter struct {
}

var DefaultGalaxyConverter = &GalaxyConverter{}

func (g *GalaxyConverter) MustFloat(input string) float64 {
	res, err := strconv.ParseFloat(input, 64)
	if err != nil {
		res = 0.0
	}
	return res
}

func (g *GalaxyConverter) MustInt(input string) int {
	rev, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}
	return rev
}

func (g *GalaxyConverter) MD5Hash(input string) string {
	hasher := md5.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}
