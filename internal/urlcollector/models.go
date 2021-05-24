package urlcollector

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

//Configuration for the Client and the Server
type Config struct {
	apiKey             string
	port               string
	url                string
	concurrentRequests int
	transportTimeout   time.Duration
	handshakeTimeout   time.Duration
	clientTimeout      time.Duration
}

func (c *Config) Setup(a string, p string, u string, cr int, tt int, ht int, ct int) {
	c.apiKey = a
	c.port = p
	c.url = u
	c.concurrentRequests = cr
	c.transportTimeout = time.Duration(tt)
	c.handshakeTimeout = time.Duration(ht)
	c.clientTimeout = time.Duration(ct)
}

//Contains all collected data
type collectedData struct {
	mutex  sync.Mutex
	Urls   []string `json:",omitempty"`
	Errors []string `json:",omitempty"`
}

func (cd *collectedData) collectError(reason string, err error) {
	e := fmt.Sprintf("with %s got error: %s", reason, err)
	cd.mutex.Lock()
	cd.Errors = append(cd.Errors, e)
	cd.mutex.Unlock()
}

func (cd *collectedData) collectURL(url string) {
	cd.mutex.Lock()
	cd.Urls = append(cd.Urls, url)
	cd.mutex.Unlock()
}

func (cd *collectedData) json() *[]byte {
	cd.mutex.Lock()
	b, err := json.Marshal(cd)
	cd.mutex.Unlock()
	if err != nil {
		cd.collectError("json.Marshal()", err)
	}
	return &b
}
