package parser

import (
	"fmt"
	"github.com/Cc-BJYG/gola/operation"
	"github.com/Cc-BJYG/gola/token"
	"strconv"
)

//Node 节点
type Node interface {
	Debug() string
	Gola() (float64, error)
}

//-----------------------------------------------------------------------------------------------------------
type numberNode struct {
	value float64
}

func (n numberNode) Debug() string {
	return "numberNode-> " + fmt.Sprint(n.value)
}
func (n numberNode) Gola() (float64, error) {
	return n.value, nil
}

//NumberNode 创建数字节点
// func NumberNode(num operation.Num) Node {
// 	n, _ := strconv.ParseFloat(num.String(), 64)
// 	return numberNode{
// 		value: n,
// 	}
// }
func NumberNode(num token.Value) Node {
	n, _ := strconv.ParseFloat(string(num), 64)
	return numberNode{
		value: n,
	}
}

// 最后好像没有用额
// //UnaryOpNode 创建一元操作符
// func UnaryOpNode(op operation.Operation, node Node) Node {
// 	return nil
// }
//-----------------------------------------------------------------------------------------------------------
type binaryOpNode struct {
	op operation.Operation

	leftNode  Node
	rightNode Node
}

func (n binaryOpNode) Debug() string {
	return "binaryOpNode-> " + "leftNode: " + n.leftNode.Debug() + " rightNode: " + n.rightNode.Debug()
}
func (n binaryOpNode) Gola() (float64, error) {
	l, err := n.leftNode.Gola()
	if err != nil {
		return 0, ErrGola
	}
	r, err := n.rightNode.Gola()
	if err != nil {
		return 0, ErrGola
	}

	return n.op(l, r)
}

//BinOpNode 创建二元操作符
func BinOpNode(op operation.Operation, leftNode Node, rightNode Node) Node {
	return binaryOpNode{
		op:        op,
		leftNode:  leftNode,
		rightNode: rightNode,
	}
}

//-----------------------------------------------------------------------------------------------------------
type buildInOpNode struct {
	op operation.Operation

	nodes []Node
}

func (n buildInOpNode) Debug() string {
	s := "binaryOpNode-> "
	for i := 0; i < len(n.nodes); i++ {
		s += fmt.Sprint("Node", i) + " " + n.nodes[i].Debug() + " "
	}
	return s
}
func (n buildInOpNode) Gola() (float64, error) {
	rs := make([]float64, 0, len(n.nodes))
	for i := 0; i < len(n.nodes); i++ {
		r, err := n.nodes[i].Gola()
		if err != nil {
			return 0, ErrGola
		}
		rs = append(rs, r)
	}
	return n.op(rs...)
}

//BuildInOpNode 创建内建函数节点
func BuildInOpNode(op operation.Operation, nodes ...Node) Node {
	return buildInOpNode{
		op:    op,
		nodes: nodes,
	}
}
