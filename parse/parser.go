package parse

import (
	"fmt"
	"github.com/pimmytrousers/pastescraper/parse/individualparsers"
)

type parserInit func() pasteParser

type pasteParser interface {
	Match(content []byte) (bool, error)
	//TODO: set this so that if we do get a sample is 100% base64 we can write it to disk decoded
	//Callback() error
}

type Parser struct {
	availableParsers map[string]pasteParser
}

var totalParsers map[string]pasteParser

func init() {
	totalParsers = map[string]pasteParser{
		"base64MZHeader":      individualparsers.Base64MZHeader{},
		"base64ELFHeader":     individualparsers.Base64ELFHeader{},
		"powershellKeyword":   individualparsers.PowershellKeyword{},
		"powershellScript":    individualparsers.PowershellScript{},
		"powershellWebClient": individualparsers.PowershellWebClient{},
		"pythonSyscall":       individualparsers.PythonSyscall{},
		"bashHeader":          individualparsers.BashBang{},
		"vbsInvocation":       individualparsers.VbsInvocation{},
	}
}

func New(specificParsers []string) (*Parser, error) {
	p := &Parser{}

	p.availableParsers = map[string]pasteParser{}

	for _, parserKey := range specificParsers {
		if _, ok := totalParsers[parserKey]; ok {
			p.availableParsers[parserKey] = totalParsers[parserKey]
		} else {
			return nil, fmt.Errorf("unknown parser type %s", parserKey)
		}
	}

	return p, nil
}

func (p *Parser) Match(content []byte) (string, error) {
	for key, parser := range p.availableParsers {
		//TODO: doesnt handle things that might match multiple signatures
		res, err := parser.Match(content)
		if err != nil {
			return key, nil
		}

		if res {
			return key, nil
		}
	}
	return "", nil
}



