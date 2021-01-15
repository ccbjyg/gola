package main

import (
	"bufio"
	"fmt"
	"github.com/Cc-BJYG/gola/parser"
	"github.com/Cc-BJYG/gola/spliter"
	"github.com/Cc-BJYG/gola/token"
	"os"
	"strings"
)

func main() {
	Parse()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')

		// convert CRLF to LF
		input = strings.Replace(input, "\r", "", -1)

		//可以设定一些内置的命令
		if strings.Compare("hi\n", input) == 0 {
			fmt.Println("hello, Yourself")
			continue
		}

		//分词
		s, err := spliter.CreateSpliter(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		tokens, err := s.Split()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//解析
		p, err := parser.CreateParser(tokens)
		if err != nil {
			fmt.Println(err)
			continue
		}
		node, err := p.Parse()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//计算
		result, err := node.Gola()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("result : ", result)

		if *debug {
			fmt.Println("tokens debug:---")
			for _, t := range tokens {
				name := token.GetTypeName(t.GetType())
				v := string(t.GetValue())
				fmt.Printf("%s %s # ", name, string(v))
			}
			fmt.Println()

			fmt.Println("node debug:---")
			fmt.Println(node.Debug())
		}
	}
}
