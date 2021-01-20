package operation

//Num 数字
type Num interface {
	String() string
}

//num
type num struct {
	// integerPart 整数部分
	// []byte 0, 1, 2, 3 ....
	//       个,十,百,千 ....
	integerPart []byte

	//dot 小数点
	dot bool

	// fractionalPart 小数部分
	// []byte 0 , 1  , 2   , 3     ....
	//        .0, .00, .000, .0000 ....
	fractionalPart []byte
}

//NewNum 创建数字
func NewNum(s []byte) (Num, error) {
	num := num{
		integerPart:    make([]byte, 0),
		dot:            false,
		fractionalPart: make([]byte, 0),
	}
	var numAppend = func(b byte) {
		num.integerPart = append([]byte{b}, num.integerPart...)
	}
	for i := 0; i < len(s); i++ {
		if s[i] <= '0' && s[i] >= '9' && s[i] != '.' {
			return num, ErrInput
		}
		if num.dot {
			if s[i] == '.' {
				return num, ErrMuchDot
			}
			numAppend(s[i])
		} else {
			if s[i] == '.' {
				num.dot = true
				numAppend = func(b byte) {
					num.fractionalPart = append(num.fractionalPart, b)
				}
				continue
			}
			numAppend(s[i])
		}
	}
	return num, nil
}

//String 输出为字符串
func (n num) String() string {
	if n.dot {
		s := make([]byte, 0, len(n.integerPart)+len(n.fractionalPart)+1)

		for i := len(n.integerPart) - 1; i >= 0; i-- {
			s = append(s, n.integerPart[i])
		}

		s = append(s, '.')

		for i := 0; i < len(n.fractionalPart); i++ {
			s = append(s, n.fractionalPart[i])
		}

		return string(s)
	}
	{
		s := make([]byte, 0, len(n.integerPart))

		for i := len(n.integerPart) - 1; i >= 0; i++ {
			s = append(s, n.integerPart[i])
		}

		return string(s)
	}
}
