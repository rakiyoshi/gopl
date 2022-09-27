package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		expr string
		want string
	}{
		{
			"sqrt(A / pi)",
			"sqrt((A / pi))",
		},
		{
			"pow(x,3) + pow(y,3)",
			"(pow(x, 3) + pow(y, 3))",
		},
		{
			"5 / 9 * (F - 32)",
			"((5 / 9) * (F - 32))",
		},
		{
			"-sqrt(-A)",
			"-sqrt(-A)",
		},
	}

	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := expr.String()
		if got != test.want {
			t.Errorf("\"%s\".String() = %q, want %q\n", test.expr, got, test.want)
		}
	}
}

func TestStringEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
	}{
		{
			"sqrt(A / pi)",
			Env{"A": 87616, "pi": math.Pi},
		},
		{
			"pow(x,3) + pow(y,3)",
			Env{"x": 12, "y": 1},
		},
		{
			"5 / 9 * (F - 32)",
			Env{"F": -40},
		},
		{
			"-sqrt(-A)",
			Env{"A": 87616},
		},
	}

	for _, test := range tests {
		expr1, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		expr2, err := Parse(expr1.String())
		if err != nil {
			t.Error(err)
		}

		value1 := fmt.Sprintf("%.6g", expr1.Eval(test.env))
		value2 := fmt.Sprintf("%.6g", expr2.Eval(test.env))
		if value1 != value2 {
			t.Errorf("\"%s\".Eval() = %q, Parse(\"%s\".String()).Eval() = Parse(\"%s\").Eval() %q\n",
				test.expr, value1, test.expr, expr1.String(), value2)
		}
	}
}
