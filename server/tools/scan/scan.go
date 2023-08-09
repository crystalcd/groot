package scan

var (
	Sf *Subfinder
	Hx *Httpx
	Nb *Naabu
)

func SetUp() {
	Sf = NewSubfinder()
	Hx = NewHttpx()
	Nb = NewNaabu()
}
