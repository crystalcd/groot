package domain

type HttpScan struct {
	Path  string
	Topic string
}

type AbstractHttpScan interface {
	Run(host string, port int, tempfile string)
}

type HttpScanUseCase interface {
	Scan(host string, ports []int) *HttpResults
}
