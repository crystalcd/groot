package tools

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/crystal/groot/global"
	wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

func Dowebanayze(jobParam Param) {
	global.G_LOG.Info("doing Dowebanayze")
	resp, err := http.DefaultClient.Get("https://www.slack.com")
	if err != nil {
		log.Fatal(err)
	}
	data, _ := ioutil.ReadAll(resp.Body) // Ignoring error for example

	wappalyzerClient, err := wappalyzer.New()
	fingerprints := wappalyzerClient.Fingerprint(resp.Header, data)
	fmt.Printf("%v\n", fingerprints)

	// Output: map[Acquia Cloud Platform:{} Amazon EC2:{} Apache:{} Cloudflare:{} Drupal:{} PHP:{} Percona:{} React:{} Varnish:{}]
}
