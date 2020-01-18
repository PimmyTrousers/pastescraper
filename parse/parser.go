package parse

import (
	"fmt"
	"github.com/pimmytrousers/pastescraper/parse/individualparsers"
)

type pasteParser interface {
	Match(content []byte) (bool, error)
	Normalize(content []byte) ([]byte, error)
}

type Parser struct {
	availableParsers map[string]pasteParser
	orderedKeys []string
}

var totalParsers map[string]pasteParser

func init() {
	totalParsers = map[string]pasteParser{
		"base64MZHeader":       	individualparsers.Base64MZHeader{},
		"base64ELFHeader":      	individualparsers.Base64ELFHeader{},
		"powershellKeyword":   		individualparsers.PowershellKeyword{},
		"powershellScript":     	individualparsers.PowershellScript{},
		"powershellWebClient":  	individualparsers.PowershellWebClient{},
		"pythonSyscall":        	individualparsers.PythonSyscall{},
		"bashHeader":           	individualparsers.BashHeader{},
		"vbsInvocation":        	individualparsers.VbsInvocation{},
		"powershellFromBase64": 	individualparsers.PowershellFromBase64{},
		"rawMZHeader": 				individualparsers.Raw64MZHeader{},
		"rawMachOHeader":			individualparsers.RawMachOHeader{},
		"largeHexBlob":				individualparsers.LargeHexBlob{},
		"base64HighEntropy":		individualparsers.Base64HighEntropy{},
		"reverseBase64MZHeader": 	individualparsers.ReverseBase64MZHeader{},
		"reverseBase64ELFHeader":	individualparsers.ReverseBase64ELFHeader{},
	}
}

func New(specificParsers []string) (*Parser, error) {
	p := &Parser{}

	p.availableParsers = map[string]pasteParser{}

	for _, parserKey := range specificParsers {
		if _, ok := totalParsers[parserKey]; ok {
			p.orderedKeys = append(p.orderedKeys, parserKey)
			p.availableParsers[parserKey] = totalParsers[parserKey]
		} else {
			return nil, fmt.Errorf("unknown parser type %s", parserKey)
		}
	}

	return p, nil
}

func (p *Parser) match(content []byte) (string, error) {
	for _, orderedKey := range p.orderedKeys {
		//TODO: doesn't handle things that might match multiple signatures
		res, err := p.availableParsers[orderedKey].Match(content)
		if err != nil {
			return "", err
		}

		if res {
			return orderedKey, nil
		}
	}
	return "", nil
}

func (p *Parser) MatchAndNormalize(content []byte) (string, []byte, error) {
	key, err := p.match(content)
	if err != nil {
		return "", nil, err
	}
	if key == "" {
		return "", nil, nil
	}

	normalizedContent, err := p.availableParsers[key].Normalize(content)

	return key, normalizedContent, nil
}





