## 设计文档

***

### 分词器

> 将输入的字符分割成数个 **token { token_type, token_value }**

#### token dict：

token type | token value 
---- | ---
NUM | **1~9** **.** 
SYMBOL | **+ - * /** ...... 
LETTER | a\~z  &&  A\~Z

### 语法分析器

> 将 **[]token** 转化为语法树

expr  -- n_expr

> 起始

n_expr -- h_expr (( **PLUS** | **MINUS** ) h_expr ) *

> 基础表达式  **+ -**  

h_expr -- vh_expr (( **MUL** |  **DIV**  ) vh_expr ) *

> 高级表达式  **\* /**

vh_expr --atom ( **CARET** atom ) *

> 乘方

atom -- **NUM**

​         -- ( **PLUS** | **MINUS** )* atom

​         -- **LPAREN**  expr  **RPAREN**

​         -- **MAX** | **MIN** | **LG** | **LN** | **FLOOR**  **LPAREN**  (expr **COMMA**) *  **RPAREN**

​         -- **ENTITY**

​         -- **ENTITY** **ASSIGN** n_expr

> 源，最小单位。同时构成逻辑闭环。



### 解析器

