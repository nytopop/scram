// eval.go

package main

import "log"

func Eval(expression Scram, e *Env) (value Scram) {
	switch exp := expression.(type) {
	case Number:
		value = exp
	case Symbol:
		value = e.Find(exp).Vars[exp]
	case []Scram:
		switch car, _ := exp[0].(Symbol); car {
		case "quote":
			value = exp[1]
		case "if":
			if Eval(exp[1], e).(bool) {
				value = Eval(exp[2], e)
			} else {
				value = Eval(exp[3], e)
			}
		case "cond":
			for _, tform := range exp[1:] {
				if Eval(tform.([]Scram)[0], e).(bool) {
					value = Eval(tform.([]Scram)[1], e)
					break
				}
			}
		case "set":
			v := exp[1].(Symbol)
			e.Find(v).Vars[v] = Eval(exp[2], e)
			value = "ok"
		case "def":
			e.Vars[exp[1].(Symbol)] = Eval(exp[2], e)
			value = "ok"
		case "lambda":
			value = Proc{exp[1], exp[2], e}
		case "begin":
			for _, i := range exp[1:] {
				value = Eval(i, e)
			}
		case "go":
			go func() { Eval(exp[1], e) }()
			value = "ok"
		default:
			operands := exp[1:]
			values := make([]Scram, len(operands))
			for i, x := range operands {
				values[i] = Eval(x, e)
			}
			value = Apply(Eval(exp[0], e), values)
		}
	default:
		log.Println("Unknown expression type - EVAL", exp)
	}
	return
}

func Apply(procedure Scram, args []Scram) (value Scram) {
	switch p := procedure.(type) {
	case func(...Scram) Scram:
		value = p(args...)
	case Proc:
		e := &Env{make(Vars), p.e}
		switch params := p.Params.(type) {
		case []Scram:
			for i, param := range params {
				e.Vars[param.(Symbol)] = args[i]
			}
		default:
			e.Vars[params.(Symbol)] = args
		}
		value = Eval(p.Body, e)
	default:
		log.Println("Unknown procedure type - APPLY", p)
	}
	return
}
