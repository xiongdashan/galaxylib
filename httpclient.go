package galaxylib

import (
	"net/http"
	"strings"
)

type GalaxyHttp struct {
}

func (g *GalaxyHttp) HeaderBulkEdit(input string, header *http.Header) {

	lines := strings.Split(input, "\n")

	//var header *http.Header
	for _, l := range lines {
		kv := strings.Split(l, ":")
		header.Add(strings.TrimSpace(kv[0]), strings.TrimSpace(kv[1]))
	}
}
