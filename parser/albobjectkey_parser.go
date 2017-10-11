package parser

import (
	"strings"
)

type ObjectKey struct {
	AccountID string
	Region    string
	ALBName   string
	ALBIP     string
	FileName  string
}

type ALBObjectKeyParser struct {
}

func NewALBObjectKeyParser() *ALBObjectKeyParser {
	return &ALBObjectKeyParser{}
}

func (p *ALBObjectKeyParser) Parse(key string) (*ObjectKey, error) {
	keys := strings.Split(key, "AWSLogs")
	if len(keys) != 2 {
		return nil, ErrSplitKey
	}
	keysufix := strings.Split(keys[1], "/")
	if len(keysufix) == 0 {
		return nil, ErrSplitKey
	}
	filename := keysufix[len(keysufix)-1]
	fkey := strings.Split(filename, "_")
	if len(fkey) < 6 {
		return nil, ErrSplitKey
	}
	return &ObjectKey{
		AccountID: fkey[0],
		Region:    fkey[2],
		ALBName:   fkey[3],
		ALBIP:     fkey[5],
		FileName:  filename,
	}, nil

}
