package operation

import (
	"math"

	"github.com/ccbjyg/gola/token"
)

var operationM = map[token.Type]Operation{
	//UnaryOperation
	token.LN:    ln,    //ln(n)
	token.LG:    lg,    //lg(n)
	token.FLOOR: floor, //floor(n)

	//BinaryOperation
	token.PLUS:  add, // 1 + 2
	token.MINUS: sub, // 1 - 2

	token.MUL: multiplication, // 1 * 2
	token.DIV: division,       // 1 ÷ 2

	token.CARET: pow, // 1 ^ 2

	//BuildInFunc
	token.MAX: max, // max(n1,n2,n3...)
	token.MIN: min, //min(n1,n2,n3...)
}

//IsUnaryOp 是否是一元操作符
func IsUnaryOp(t token.Token) (Operation, bool) {
	if t == nil {
		return nil, false
	}
	op, exist := unaryOperationM[t.GetType()]
	return op, exist
}

//IsBinaryOp 是否是二元操作符
func IsBinaryOp(t token.Token) (Operation, bool) {
	if t == nil {
		return nil, false
	}
	op, exist := binaryOperationM[t.GetType()]
	return op, exist
}

//IsBuildInOp 是否内建函数
func IsBuildInOp(t token.Token) (Operation, bool) {
	if t == nil {
		return nil, false
	}
	op, exist := buildInOperationM[t.GetType()]
	return op, exist
}

//UnaryOperation 一元操作符 -1
var unaryOperationM = map[token.Type]Operation{
	token.PLUS:  add, //+
	token.MINUS: sub, //-
}

//BinaryOperation
var binaryOperationM = map[token.Type]Operation{
	// 1 op 2
	token.PLUS:  add, //+
	token.MINUS: sub, //-

	token.MUL: multiplication, //*
	token.DIV: division,       //÷

	token.CARET: pow, //^
}

//BuildInFunc 内建函数
var buildInOperationM = map[token.Type]Operation{
	token.LN:    ln,    //ln(n)
	token.LG:    lg,    //lg(n)
	token.FLOOR: floor, //floor(n)

	token.MAX: max, // max(n1,n2,n3...)
	token.MIN: min, //min(n1,n2,n3...)
}

//Operation 操作符
type Operation func(...float64) (float64, error)

//base operation
func add(nums ...float64) (float64, error) {
	if len(nums) == 1 {
		return nums[0], nil
	}
	if len(nums) == 2 {
		return nums[0] + nums[1], nil
	}
	return 0, ErrInput
}
func sub(nums ...float64) (float64, error) {
	if len(nums) == 1 {
		return 0 - nums[0], nil
	}
	if len(nums) == 2 {
		return nums[0] - nums[1], nil
	}
	return 0, ErrInput
}
func multiplication(nums ...float64) (float64, error) {
	if len(nums) != 2 {
		return 0, ErrInput
	}
	return nums[0] * nums[1], nil
}
func division(nums ...float64) (float64, error) {
	if len(nums) != 2 {
		return 0, ErrInput
	}
	if nums[1] == 0 {
		return 0, ErrInput
	}
	return nums[0] / nums[1], nil
}
func pow(nums ...float64) (float64, error) {
	if len(nums) != 2 {
		return 0, ErrInput
	}
	return math.Pow(nums[0], nums[1]), nil
}

//build-in function

func max(nums ...float64) (float64, error) {
	if len(nums) < 1 {
		return 0, ErrInput
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max, nil
}
func min(nums ...float64) (float64, error) {
	if len(nums) < 1 {
		return 0, ErrInput
	}
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min, nil
}
func ln(nums ...float64) (float64, error) {
	if len(nums) != 1 {
		return 0, ErrInput
	}
	return math.Logb(nums[0]), nil
}
func lg(nums ...float64) (float64, error) {
	if len(nums) != 1 {
		return 0, ErrInput
	}
	return math.Log10(nums[0]), nil
}
func floor(nums ...float64) (float64, error) {
	if len(nums) != 1 {
		return 0, ErrInput
	}
	return math.Floor(nums[0]), nil
}
