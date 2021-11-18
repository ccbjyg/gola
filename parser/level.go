package parser

import "github.com/ccbjyg/gola/token"

var nLevel = map[token.Type]bool{
	token.PLUS:  true,
	token.MINUS: true,
}

var hLevel = map[token.Type]bool{
	token.MUL: true,
	token.DIV: true,
}
var vhLevel = map[token.Type]bool{
	token.CARET: true,
}
