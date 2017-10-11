package parser

import (
	"errors"
)

var (
	// ErrIgnored should be set, If parsed line should be ignored.
	ErrIgnored   = errors.New("line ignored")
	ErrNoSupport = errors.New("not supported")
	ErrRegexp    = errors.New("regexp faild")
	ErrSplitKey  = errors.New("split key faild")
)
