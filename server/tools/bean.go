package tools


type Config struct {
	Path string
}

type Param struct {
	Target string
	Domain string
}

type Domain struct {
	Project string `bson:"project,omitempty"`
	Domain string `bson:"domain,omitempty"`
	From string `bson:"from,omitempty"`
	
}

