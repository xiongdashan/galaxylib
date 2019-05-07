package galaxylib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type GalaxyLokiClient struct {
	API    string
	Status bool
	Label  string
}

type LokiData struct {
	Streams []*LokiStream `json:"streams"`
}

type LokiStream struct {
	Entries []*LokiEntry `json:"entries"`
	Labels  string       `json:"labels"`
}

type LokiEntry struct {
	Line string    `json:"line"`
	Ts   time.Time `json:"ts"`
}

func NewGalaxyLokiClient(api string, status bool, label string) *GalaxyLokiClient {
	return &GalaxyLokiClient{
		API:    api,
		Status: status,
		Label:  label,
	}
}

func DefaultGalaxyLokiClient() *GalaxyLokiClient {
	api := GalaxyCfgFile.MustValue("loki", "api")
	status := GalaxyCfgFile.MustBool("loki", "status")
	label := GalaxyCfgFile.MustValue("loki", "label")
	return NewGalaxyLokiClient(api, status, label)
}

func (l *GalaxyLokiClient) Send(name string, data interface{}) {
	if l.Status == false {
		return
	}
	l.sender(name, data)
}

func (l *GalaxyLokiClient) SendAsync(name string, data interface{}) {
	if l.Status == false {
		return
	}
	go l.sender(name, data)
}

func (l *GalaxyLokiClient) sender(name string, data interface{}) {

	lables := fmt.Sprintf(`{%s="%s"}`, l.Label, name)
	var entries []*LokiEntry

	buf, err := json.Marshal(data)
	if err != nil {
		GalaxyLogger.Error(err.Error())
		return
	}

	entries = append(entries, &LokiEntry{
		Line: string(buf),
		Ts:   time.Now(),
	})

	streams := append([]*LokiStream{}, &LokiStream{
		Labels:  lables,
		Entries: entries,
	})

	sendData := new(LokiData)
	sendData.Streams = streams

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(sendData); err != nil {
		GalaxyLogger.Error(err.Error())
		return
	}

	//fmt.Println(b.String())

	rs, err := http.Post(l.API, "application/json", b)
	if err != nil {
		GalaxyLogger.Error(err.Error())
	}
	defer rs.Body.Close()
	rsBuf, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("loki return: %s\n", string(rsBuf))
}
