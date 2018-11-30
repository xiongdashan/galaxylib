package galaxylib

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/text/encoding/simplifiedchinese"
)

type GalaxyTools struct {
}

var DefaultGalaxyTools = &GalaxyTools{}

//func(t *GalaxyTools) Init
func (t *GalaxyTools) Bytes2CHString(buf []byte) string {
	decode, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(buf)
	//simplifiedchinese.GB18030.NewEncoder().
	return string(decode)
}

func (t *GalaxyTools) Contains(ary []string, item string) bool {

	if len(ary) == 0 {
		return false
	}

	for _, val := range ary {
		if val == item {
			return true
		}
	}
	return false
}

func (t *GalaxyTools) ResponseToString(coler io.ReadCloser) string {
	buf, _ := ioutil.ReadAll(coler)
	return string(buf)
}

func (t *GalaxyTools) ReadFileLine(name string, fn func(line string)) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fn(scanner.Text())
	}
}
