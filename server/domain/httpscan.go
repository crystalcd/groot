package domain

type HttpScan struct {
	Path  string
	Topic string
}

type AbstractHttpScan interface {
	Run(host string, port string, tempfile string)
}

type HttpScanUseCase interface {
	Scan(host string, ports []string) *HttpResults
}
