package scan

import (
	"io/ioutil"
	"net/http"

	wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

type Wappalyze struct {
}

func NewWappalyze() *Wappalyze {
	return &Wappalyze{}
}

func (w *Wappalyze) Scan(url string) (map[string]struct{}, error) {
	rs := make(map[string]struct{})
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return rs, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return rs, err
	}
	wappalyerClient, err := wappalyzer.New()
	if err != nil {
		return rs, err
	}
	rs = wappalyerClient.Fingerprint(resp.Header, data)
	return rs, nil
}
