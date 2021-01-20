package operation

import "errors"

var (
	//ErrInput 输入错误
	ErrInput = errors.New("input error")

	//ErrMuchDot 小数点过多
	ErrMuchDot = errors.New("too many dot")
)
