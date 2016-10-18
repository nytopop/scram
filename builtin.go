// builtin.go

package main

import "reflect"

var globalenv Env

func init() {
	globalenv = Env{
		Vars{
			// List manipulation
			"cons": func(a ...Scram) Scram {
				switch car := a[0]; cdr := a[1].(type) {
				case []Scram:
					return append([]Scram{car}, cdr...)
				default:
					return []Scram{car, cdr}
				}
			},
			"car": func(a ...Scram) Scram {
				return a[0].([]Scram)[0]
			},
			"cdr": func(a ...Scram) Scram {
				return a[0].([]Scram)[1:]
			},
			//"list": Eval(read(
			//	"(lambda z z)"),
			//	&globalenv),
			// Concurrency
			"chan": func(a ...Scram) Scram {
				return make(chan Scram, int(a[0].(Number)))
			},
			"->": func(a ...Scram) Scram {
				a[0].(chan Scram) <- a[1]
				return "ok"
			},
			"<-": func(a ...Scram) Scram {
				return <-a[0].(chan Scram)
			},
			// Arithmetic
			"+": func(a ...Scram) Scram {
				v := a[0].(Number)
				for _, i := range a[1:] {
					v += i.(Number)
				}
				return v
			},
			"-": func(a ...Scram) Scram {
				v := a[0].(Number)
				for _, i := range a[1:] {
					v -= i.(Number)
				}
				return v
			},
			"*": func(a ...Scram) Scram {
				v := a[0].(Number)
				for _, i := range a[1:] {
					v *= i.(Number)
				}
				return v
			},
			"/": func(a ...Scram) Scram {
				v := a[0].(Number)
				for _, i := range a[1:] {
					v /= i.(Number)
				}
				return v
			},
			// Comparison
			"eq": func(a ...Scram) Scram {
				return reflect.DeepEqual(a[0], a[1])
			},
			"lt": func(a ...Scram) Scram {
				return a[0].(Number) < a[1].(Number)
			},
			"lte": func(a ...Scram) Scram {
				return a[0].(Number) <= a[1].(Number)
			},
			"gt": func(a ...Scram) Scram {
				return a[0].(Number) > a[1].(Number)
			},
			"gte": func(a ...Scram) Scram {
				return a[0].(Number) >= a[1].(Number)
			},
		}, nil}
}
