package galaxylib

import (
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type Jar struct {
	lk      sync.Mutex
	cookies map[string][]*http.Cookie
}

func NewJar() *Jar {
	jar := new(Jar)
	jar.cookies = make(map[string][]*http.Cookie)
	return jar
}

// SetCookies handles the receipt of the cookies in a reply for the
// given URL.  It may or may not choose to save the cookies, depending
// on the jar's policy and implementation.
func (jar *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	jar.lk.Lock()
	jar.cookies[u.Host] = cookies
	jar.lk.Unlock()
}

// Cookies returns the cookies to send in a request for the given URL.
// It is up to the implementation to honor the standard cookie use
// restrictions such as in RFC 6265.
func (jar *Jar) Cookies(u *url.URL) []*http.Cookie {
	return jar.cookies[u.Host]
}

func (j *Jar) BulkEdit(cookieView string) []*http.Cookie {
	ary := strings.Split(cookieView, ";")
	var cookies []*http.Cookie
	for _, v := range ary {
		keyVal := strings.Split(v, "=")
		c := &http.Cookie{
			Name:  strings.TrimSpace(keyVal[0]),
			Value: strings.TrimSpace(keyVal[1]),
		}
		cookies = append(cookies, c)
	}
	return cookies
}
