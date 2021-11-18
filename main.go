package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ccbjyg/gola/parser"
	"github.com/ccbjyg/gola/spliter"
	"github.com/ccbjyg/gola/token"
)

func main() {
	Parse()

	//运行文件，未来需要加更多的调试信息，将这部分内容精细化。
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

		//分词--------------------------------------------------------------------------------------------------------
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
		if *debug {
			fmt.Println("tokens debug:---")
			for _, t := range tokens {
				name := token.GetTypeName(t.GetType())
				v := string(t.GetValue())
				fmt.Printf("%s %s # ", name, string(v))
			}
			fmt.Println()
		}

		//解析--------------------------------------------------------------------------------------------------------
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
		if *debug {
			//解析器debug
		}

		//计算--------------------------------------------------------------------------------------------------------
		if node == nil {
			continue
		}
		result, err := node.Gola()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("result : ", result)
	}
}
