package usage_parser

type UsageParser interface {
	Parse(input []string) (result []*UsageResult)
}

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(input []string) (result []*UsageResult) {
	return []*UsageResult{
		{},
	}
}
