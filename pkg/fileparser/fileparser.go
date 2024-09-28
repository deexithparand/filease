package fileparser

type Parser interface {
	Parse(filepath string) (string, error)
}

func ParseFile(p Parser, filepath string) (string, error) {
	return p.Parse(filepath)
}
