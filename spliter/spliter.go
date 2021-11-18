package spliter

import (
	"errors"

	"github.com/ccbjyg/gola/token"
)

//Spliter 分词器
type Spliter interface {
	Split() ([]token.Token, error)
}

//
type spliter struct {
	pos    int    //当前位置
	source []byte //源公式
	//用于放置所有对象
	tokens []token.Token
}

//CreateSpliter 新建分词器
func CreateSpliter(input string) (Spliter, error) {
	inputb := []byte(input)
	s := &spliter{
		pos:    0,
		source: inputb,
		tokens: make([]token.Token, 0),
	}
	return s, nil
}

func (s *spliter) Split() ([]token.Token, error) {
	err := s.split()
	return s.tokens, err
}

//split 将输入的
func (s *spliter) split() error {
	for s.pos < len(s.source) {
		b := s.source[s.pos]

		//几种非常特殊的情况
		if b == '\n' || b == '\t' {
			return nil
		}
		if b == ' ' {
			s.pos++
			continue
		}
		//声明token
		var token token.Token
		var err error

		//判断起始点
		if isNums(b) {
			token, err = s.makeNum()
		} else if isSymbol(b) {
			token, err = s.makeSymbol()
		} else if isLetter(b) {
			token, err = s.makeString()
		} else {
			return errors.New("get token error")
		}

		//添加token
		s.tokens = append(s.tokens, token)
		if err != nil {
			return err
		}
	}
	return nil
}

//用于生成连续的数字 token
func (s *spliter) makeNum() (token.Token, error) {
	num := []byte{}
	hasDot := false

	//lookup
	for s.pos < len(s.source) {
		b := s.source[s.pos]

		if isNums(b) {
			s.pos++
			num = append(num, b)
		} else if b == '.' {
			s.pos++
			num = append(num, b)
			if hasDot {
				//不管怎样， "12.3.?" 都是非法的，后面的就不管了。
				return token.CreateToken(token.UNKNOWN, num), ErrInput
			}
			hasDot = true
		} else {
			//其它情况返回数字
			//返回前判断末尾是否合规
			if num[len(num)-1] == '.' {
				return token.CreateToken(token.UNKNOWN, num), ErrInput
			}
			return token.CreateToken(token.NUM, num), nil
		}
	}
	return token.CreateToken(token.UNKNOWN, num), ErrInput
}

//用于生成符号 token
func (s *spliter) makeSymbol() (token.Token, error) {
	// 最长匹配原则
	// 基于如下事实：
	//   复合符号可以拆分为多个字符号。
	//   尽可能匹配更长的符号。
	symbol := []byte{}

	//lookup
	for s.pos < len(s.source) {
		b := s.source[s.pos]

		if isSymbol(b) {
			s.pos++
			symbol = append(symbol, b)
			//不是复合符号就继续贪心匹配即可
			if len(symbol) < 2 {
				continue
			}
			//检查是否满足更长的复合符号
			symbolType := token.Match(token.SYMBOL, symbol...)
			if symbolType == token.UNKNOWN {
				s.pos--
				symbol = symbol[:len(symbol)-1]
				symbolType = token.Match(token.SYMBOL, symbol...)
				return token.CreateToken(symbolType, symbol), nil
			}
		} else {
			symbolType := token.Match(token.SYMBOL, symbol...)
			return token.CreateToken(symbolType, symbol), nil
		}
	}
	return token.CreateToken(token.UNKNOWN, symbol), ErrInput
}

//用于生成字符串 token
func (s *spliter) makeString() (token.Token, error) {
	//大致思路和符号匹配类似，但似乎并不需要回滚。
	str := []byte{}

	//lookup
	for s.pos < len(s.source) {
		b := s.source[s.pos]

		//允许 a1 这种定义方式，不允许 1a
		if isLetter(b) || isNums(b) {
			s.pos++
			str = append(str, b)
		} else {
			wordType := token.Match(token.KEYWORDS, str...)
			return token.CreateToken(wordType, str), nil
		}
	}
	return token.CreateToken(token.UNKNOWN, str), ErrInput
}

//ASCII 标准字符

//isNums 是否为数字
func isNums(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

//isSymbol 是否为基本符号(不确定这些基础符号是否是连续的...)
func isSymbol(b byte) bool {
	//因为是 ASCII 标准字符，所以也可以用其他方式匹配。
	if tt := token.Match(token.SYMBOL, b); tt != token.UNKNOWN {
		return true
	}
	return false
}

//isLetter 是否为字母
func isLetter(b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
		return true
	}
	return false
}
