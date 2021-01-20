package token

//Kind token 类别，用于Match
type Kind int

const (
	//SYMBOL 符号
	SYMBOL Kind = iota

	//KEYWORDS 关键词
	KEYWORDS
)

//Type token类型
type Type int

//Value 值
type Value []byte

const (
	//UNKNOWN 未知
	UNKNOWN Type = iota

	//计划:
	//基本数据类型: number, string, bool
	//高级数据类型: arrary, struct, func

	//NUM 整数 小数...
	NUM
	//ENTITY 对象 (本意就是指一段 byte数组了 -_-)
	ENTITY

	//symbol

	symbolStart

	//PLUS +
	PLUS
	//MINUS -
	MINUS
	//MUL 乘 *
	MUL
	//DIV 除 /
	DIV
	//CARET 乘方 ^
	CARET

	//ASSIGN =
	ASSIGN

	//LPAREN (
	LPAREN

	// //LBRACK [
	// LBRACK

	// //LBRACE {
	// LBRACE

	//COMMA ,
	COMMA

	// //DOT .
	// DOT

	//RPAREN )
	RPAREN
	// //RBRACK ]
	// RBRACK
	// //RBRACE }
	// RBRACE
	// //SEMICOLON ;
	// SEMICOLON
	// //COLON :
	// COLON

	symbolEnd

	//keywords
	keywordsStart
	//内置运算符

	//MAX 取最大值
	MAX
	//MIN 取最小值
	MIN
	//LN log以e为底的对数
	LN
	//LG log以10为底的对数
	LG
	//FLOOR 舍小数
	FLOOR

	keywordsEnd
)

//token 对应的名字，调试用
var tokens = [...]string{
	UNKNOWN: "unknown",
	//NUM 整数 小数...
	NUM: "NUM",
	//ENTITY 对象
	ENTITY: "ENTITY",

	//Base Operation 普通操作

	//PLUS +
	PLUS: "+",
	//MINUS -
	MINUS: "-",
	//MUL 乘 *
	MUL: "*",
	//DIV 除 /
	DIV: "/",
	//CARET 乘方 ^
	CARET: "^",

	//ASSIGN =
	ASSIGN: "=",

	//LPAREN (
	LPAREN: "(",
	// //LBRACK [
	// LBRACK: "[",
	// //LBRACE {
	// LBRACE: "{",
	//COMMA ,
	COMMA: ",",
	// //DOT .
	// DOT: ".",

	//RPAREN )
	RPAREN: ")",
	// //RBRACK ]
	// RBRACK: "]",
	// //RBRACE }
	// RBRACE: "}",
	// //SEMICOLON ;
	// SEMICOLON: ";",
	// //COLON :
	// COLON: ":",

	//keywords

	//内置运算符selfFunc

	//MAX 取最大值
	MAX: "max",
	//MIN 取最小值
	MIN: "min",
	//LN log以e为底的对数
	LN: "ln",
	//LG log以10为底的对数
	LG: "lg",
	//FLOOR 舍小数
	FLOOR: "floor",
}

//GetTypeName 获取Token 的名字
func GetTypeName(t Type) string {
	return tokens[t]
}

func init() {
	for i := symbolStart + 1; i < symbolEnd; i++ {
		symbolM[tokens[i]] = i
	}
	for i := keywordsStart + 1; i < keywordsEnd; i++ {
		keyWords[tokens[i]] = i
	}
}

//Match 匹配类型
func Match(kind Kind, bs ...byte) Type {
	if kind == SYMBOL {
		return getSymbolType(bs...)
	}
	if kind == KEYWORDS {
		return getKeywordsType(bs...)
	}
	return UNKNOWN
}

// var nums = map[byte]Type{'0': NUM, '1': NUM, '2': NUM, '3': NUM, '4': NUM, '5': NUM, '6': NUM, '7': NUM, '8': NUM, '9': NUM}

//关键符号表，需要和上面对应
var symbolM = map[string]Type{}

func getSymbolType(bs ...byte) Type {
	if tt, exist := symbolM[string(bs)]; exist {
		return tt
	}
	return UNKNOWN
}

//关键字表
var keyWords = map[string]Type{}

func getKeywordsType(bs ...byte) Type {
	//匹配到关键字则是关键字，否则作为 变量实体返回
	if tt, exist := keyWords[string(bs)]; exist {
		return tt
	}
	return ENTITY
}
