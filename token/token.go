package token

//Token 子段
type Token interface {
	GetType() Type
	GetValue() Value
}

//CreateToken 创建Token
func CreateToken(tt Type, tv []byte) Token {
	return token{
		tt: tt,
		tv: tv,
	}
}

//token 实现
type token struct {
	tt Type
	tv Value
}

func (t token) GetType() Type {
	return t.tt
}
func (t token) GetValue() Value {
	return t.tv
}
