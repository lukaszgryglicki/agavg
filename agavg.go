package main

import (
	"fmt"
	"math/cmplx"
	"os"
	"strconv"
	"strings"
)

func fatalf(f string, args ...interface{}) {
	fmt.Printf(f, args...)
	os.Exit(1)
}

func getComplex(str string) (c complex128, err error) {
	str = strings.Replace(strings.ToLower(str), "-", "+-", -1)
	if str[:1] == "+" {
		str = str[1:]
	}
	var f float64
	re := 0.0
	im := 0.0
	ary := strings.Split(str, "+")
	for _, it := range ary {
		it = strings.TrimSpace(it)
		r := true
		if strings.HasSuffix(it, "i") {
			r = false
			it = it[:len(it)-1]
			if it == "" {
				it = "1"
			} else if it == "-" {
				it = "-1"
			}
		}
		f, err = strconv.ParseFloat(it, 64)
		if err != nil {
			return
		}
		if r {
			re += f
		} else {
			im += f
		}
	}
	c = complex(re, im)
	return
}

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s '1.2+3.1i' '-0.5-1.12i'\n", os.Args[0])
		return
	}
	dbg := false
	if os.Getenv("DBG") != "" {
		dbg = true
	}
	args := []complex128{}
	for _, s := range os.Args[1:] {
		x, err := getComplex(s)
		if err != nil {
			fatalf("%s: %+v\n", s, err)
		}
		args = append(args, x)
	}
	n := len(args)
	invN := complex(1.0/float64(n), 0.0)
	i := 0
	for {
		a := complex(0.0, 0.0)
		g := complex(1.0, 0.0)
		for _, x := range args {
			a += x
			g *= x
		}
		na := a * invN
		ng := cmplx.Pow(g, invN)
		d := cmplx.Abs(na - ng)
		if dbg {
			fmt.Printf("%+v --> (%v,%v) [%f]\n", args, na, ng, d)
		}
		args = []complex128{na, ng}
		if n != 2 {
			n = 2
			invN = complex(0.5, 0.0)
		}
		i++
		if d < 1e-13 || i == 100 {
			break
		}
	}
	fmt.Printf("%v (in %d steps)\n", args[0], i)
}
