package parser

import (
	"fmt"
	"strconv"

	"github.com/ccbjyg/gola/operation"
	"github.com/ccbjyg/gola/token"
)

//Node 节点
type Node interface {
	Debug() string
	Gola() (float64, error)
}

//-----------------------------------------------------------------------------------------------------------
//数字节点
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
func NumberNode(num token.Value) Node {
	n, _ := strconv.ParseFloat(string(num), 64)
	return numberNode{
		value: n,
	}
}

//-----------------------------------------------------------------------------------------------------------
//字符串节点
type stringNode struct {
	value []byte
}

func (n stringNode) Debug() string {
	return "stringNode-> " + fmt.Sprint(n.value)
}
func (n stringNode) Gola() (float64, error) {
	return 0, nil
}

//StringNode 创建数字节点
func StringNode(str token.Value) Node {
	return stringNode{
		value: str,
	}
}

//-----------------------------------------------------------------------------------------------------------
//变量节点，变量节点掌握着另一个node的头部。
type varNode struct {
	node Node
}

func (n varNode) Debug() string {
	return "varNode-> " + n.node.Debug()
}
func (n varNode) Gola() (float64, error) {
	return n.node.Gola()
}

//VarNode 创建数字节点
func VarNode(node Node) Node {
	return varNode{
		node: node,
	}
}

//-----------------------------------------------------------------------------------------------------------
//实际上 operation 那边也有相关的容错逻辑，以下几个操作符事实上都可以合并的 -_-
//单目操作节点、双目操作节点、内置函数节点
type operationNode struct {
	op operation.Operation

	nodes []Node
}

func (n operationNode) Debug() string {
	s := "binaryOpNode-> "
	for i := 0; i < len(n.nodes); i++ {
		s += fmt.Sprint("Node", i) + " " + n.nodes[i].Debug() + " "
	}
	return s
}
func (n operationNode) Gola() (float64, error) {
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

//OperationNode 创建内建函数节点
func OperationNode(op operation.Operation, nodes ...Node) Node {
	return operationNode{
		op:    op,
		nodes: nodes,
	}
}
