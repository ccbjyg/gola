package parser

import "errors"

var (
	//ErrInput 输入错误
	ErrInput = errors.New("input error")

	//ErrGola 运行错误
	ErrGola = errors.New("node calculate error")
)
