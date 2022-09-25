package eval

import (
	"fmt"
	"strconv"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'f', -1, 64)
}

func (u unary) String() string {
	return fmt.Sprintf("%s%s", string(u.op), u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %s %s)", b.x.String(), string(b.op), b.y.String())
}

func (c call) String() string {
	var argsStr string
	for _, arg := range c.args {
		if len(argsStr) == 0 {
			argsStr += arg.String()
		} else {
			argsStr += ", " + arg.String()
		}
	}
	return fmt.Sprintf("%s(%s)", c.fn, argsStr)
}
