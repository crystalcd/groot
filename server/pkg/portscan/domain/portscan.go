package domain

type PortScan struct {
	Path  string
	Topic string
}

type AbstractPortScan interface {
	Run(domain string, tempfile string)
}

type PortScanUseCase interface {
	Scan(target string) *Result
}
