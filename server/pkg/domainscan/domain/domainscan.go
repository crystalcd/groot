package domain

type Config struct {
	Path string
}

type DomainScan struct {
	Config Config
	Topic  string
}

type AbstractDomainScan interface {
	Run(domain string, tempfile string)
}

type DomainScanUseCase interface {
	Scan(target string) *Result
}
