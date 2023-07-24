package domain

type HttpScan struct {
	Path  string
	Topic string
}

type AbstractHttpScan interface {
	Run(domain string, port int, tempfile string)
}

type HttpScanUseCase interface {
	Scan(target string) *Result
}
