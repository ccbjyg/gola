package parser

import "errors"

//定义global 变量

var globalEntity = map[string]Node{}

func registerGloblEntity(name string, node Node) error {
	globalEntity[name] = node
	return nil
}

func getGlobalEntity(name string) (Node, error) {
	if node, exist := globalEntity[name]; exist {
		return node, nil
	}
	return nil, errors.New("undefined " + name)
}
