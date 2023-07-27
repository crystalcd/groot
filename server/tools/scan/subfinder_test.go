package scan_test

import (
	"log"
	"testing"

	"github.com/crystal/groot/tools/scan"
)

func TestScan(t *testing.T) {
	subfinder := scan.NewSubfinder()
	subdomains, err := subfinder.Scan("baidu.com")
	if err != nil {
		log.Panic(err)
	}
	log.Printf("subdomains: %v", subdomains)
}

func TestNaabu(t *testing.T) {
	naabu := scan.NewNaabu()
	ports, err := naabu.Scan("baidu.com")
	if err != nil {
		log.Panic(err)
	}
	log.Printf("ports: %v", ports)
}
