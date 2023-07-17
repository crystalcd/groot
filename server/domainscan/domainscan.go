package domainscan

type DomainScan interface {
	Do()
	AsyncDo()
	Run(domain string)
	ParseResult(domain string, data []byte)
	Write2MongoDB()
}
