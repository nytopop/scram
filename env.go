// env.go

package main

type Proc struct {
	Params, Body Scram
	e            *Env
}

type Vars map[Symbol]Scram
type Env struct {
	Vars
	Outer *Env
}

func (e *Env) Find(s Symbol) *Env {
	if _, ok := e.Vars[s]; ok {
		return e
	} else {
		return e.Outer.Find(s)
	}
}
