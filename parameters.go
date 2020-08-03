package dota2api

type Parameter interface {
	key() string
	value() interface{}
}

type ParameterInt struct {
	k string
	v int
}

func (p ParameterInt) key() string {
	return p.k
}

func (p ParameterInt) value() int {
	return p.v
}

type ParameterInt64 struct {
	k string
	v int64
}

func (p ParameterInt64) key() string {
	return p.k
}

func (p ParameterInt64) value() int64 {
	return p.v
}
