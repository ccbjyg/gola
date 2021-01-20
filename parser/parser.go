package parser

import (
	"errors"

	"github.com/Cc-BJYG/gola/operation"
	"github.com/Cc-BJYG/gola/token"
)

//Parser 语法解析器
type Parser interface {
	Parse() (Node, error)
}
type parser struct {
	root   Node
	pos    int
	tokens []token.Token
}

//CreateParser 创建解析器
func CreateParser(tokens []token.Token) (Parser, error) {
	p := &parser{
		pos:    0,
		tokens: tokens,
	}
	return p, nil
}

func (p *parser) Parse() (Node, error) {
	root, err := p.expr()
	p.root = root
	return root, err
}

//getCurrentToken 用于获取下一个token
func (p *parser) getCurrentToken() token.Token {
	curpos := p.pos
	if curpos < len(p.tokens) {
		p.pos++
		return p.tokens[curpos]
	}
	return nil
}

//解析过程可以视为从上至下

func (p *parser) expr() (Node, error) {
	// expr -- nExpr
	return p.nExpr()
}

//+ - 属于二元操作符
func (p *parser) nExpr() (Node, error) {
	// nExpr -- hExpr (( PLUS | MINUS ) hExpr ) *
	return p.binaryOpNode(p.hExpr, nLevel)
}

// * ÷
func (p *parser) hExpr() (Node, error) {
	// hExpr -- vhExpr (( PLUS | MINUS ) vhExpr ) *
	return p.binaryOpNode(p.vhExpr, hLevel)
}

// ^
func (p *parser) vhExpr() (Node, error) {
	return p.binaryOpNode(p.atom, vhLevel)
}

//atom 数字源，同时用于构成闭环。
func (p *parser) atom() (Node, error) {
	//起始可以是 数字 ( 一元操作符 内建函数
	t := p.getCurrentToken()

	if t != nil && t.GetType() == token.NUM {
		// NUM 返回数字节点
		// num, err := operation.NewNum(t.GetValue())
		// return NumberNode(num), err
		return NumberNode(t.GetValue()), nil
	} else if t != nil && t.GetType() == token.LPAREN {
		// ( nExpr ) 括号
		// expr, err := p.nExpr()
		node, err := p.expr()
		if err != nil {
			return node, err
		}

		t = p.getCurrentToken() //这里在迭代后需要再更新一下
		if t == nil || t.GetType() != token.RPAREN {
			return node, errors.New("miss ')' ")
		}

		return node, nil
	} else if t != nil && t.GetType() == token.ENTITY {
		//支持变量
		entityName := string(t.GetValue()) // 获取变量名

		node, _ := getGlobalEntity(entityName)
		if node != nil {
			//说明已经注册过了
			return node, nil
		}
		//否则通过 = 来判断进行注册

		//判断是否为 =
		t = p.getCurrentToken()
		if t == nil || t.GetType() != token.ASSIGN {
			return nil, errors.New(entityName + " miss = symbol")
		}

		//解析 node
		node, err := p.nExpr()
		if err != nil {
			return nil, err
		}

		//注册全局 node
		return nil, registerGloblEntity(string(entityName), node)
	} else if op, ok := operation.IsUnaryOp(t); ok {
		//支持 +++1 ---1
		node, err := p.atom()
		if err != nil {
			return node, err
		}
		return UnaryOpNode(op, node), nil
	} else if op, ok := operation.IsBuildInOp(t); ok {
		//如果是内置函数
		t = p.getCurrentToken()
		if t == nil || t.GetType() != token.LPAREN {
			return BuildInOpNode(op), errors.New("miss ( ")
		}

		nodes := []Node{}
		lookup := true
		for lookup {
			node, err := p.expr()
			if err != nil {
				return node, err
			}
			nodes = append(nodes, node)
			//获取下一个 看是否为 ','
			t = p.getCurrentToken()
			if t == nil || t.GetType() != token.COMMA {
				//如果 下一个token为空或者 不为 COMMA ','
				//停止查找并回滚
				lookup = false
				p.pos--
			}
		}

		t = p.getCurrentToken()
		if t == nil || t.GetType() != token.RPAREN {
			return BuildInOpNode(op, nodes...), errors.New("miss ) ")
		}
		return BuildInOpNode(op, nodes...), nil
	} else {
		//都没开始就结束了
		if t != nil {
			return nil, errors.New("error token: " + token.GetTypeName(t.GetType()) + " " + string(t.GetValue()))
		}
		return nil, errors.New("miss start symbol num ,keywords, ( and so on")
	}
}

//根据一定规则返回一个二元操作节点
func (p *parser) binaryOpNode(f func() (Node, error), tokenTypes map[token.Type]bool) (Node, error) {
	leftNode, err := f()
	if err != nil {
		return leftNode, err
	}

	lookup := true
	for lookup {
		t := p.getCurrentToken()

		if t != nil && tokenTypes[t.GetType()] {
			op, ok := operation.IsBinaryOp(t) //获取该操作符
			if !ok {
				return leftNode, errors.New("error operation")
			}
			rightNode, err := f()
			if err != nil {
				return BinOpNode(op, leftNode, rightNode), err
			}
			//左结合
			leftNode = BinOpNode(op, leftNode, rightNode)
		} else {
			//回滚
			lookup = false
			p.pos--
		}

	}
	return leftNode, nil
}
