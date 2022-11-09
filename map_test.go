package main

import "testing"

type context interface {
	Name()
}
type context2 struct {
}

func (c context2) Name() {

}

var m map[string]func(c context) = make(map[string]func(c context))

func signUp(c context) {
	c.Name()
}

func Test1(t *testing.T) {
	m["hello"] = signUp
}
