package main

import (
	"bufio"
	"fmt"
	"gopl/ch07/eval"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()
	if s == "" {
		fmt.Fprintf(os.Stderr, "empty expression")
		os.Exit(1)
	}
	expr, err := eval.Parse(s)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	env := make(eval.Env)
	for name := range vars {
		fmt.Printf("%s = ", name)
		sc.Scan()
		input := sc.Text()
		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		env[name] = value
	}

	fmt.Printf("%.6g\n", expr.Eval(env))
}
