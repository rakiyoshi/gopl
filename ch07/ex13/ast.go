package eval

// Expr is the arithmetic expression
type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
	String() string
}

// Var identifies variable. e.g. x
type Var string

// literal is numeric const value. e.g. 3.141
type literal float64

// unary is single arithmetic operation. e.g. -x
type unary struct {
	op rune // '+', or '-'
	x  Expr
}

// binary is binary arithmetic operation. e.g. x+y
type binary struct {
	op   rune // '+', '-', '*', or '/'
	x, y Expr
}

// call is function callee
type call struct {
	fn   string // "pow", "sin", or "sqrt"
	args []Expr
}
