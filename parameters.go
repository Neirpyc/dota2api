package dota2api

import "errors"

const (
	parameterKindHeroId = iota
	parameterKindMatchesRequested
	parameterKindAccountId
	parameterKindStartAtMatchId
	parameterKindMinPlayers
	parameterKindMatchId
	parameterStartMatchAtSeqNum
)

type Parameter interface {
	key() string
	value() interface{}
	kind() int
}

func getParameterMap(require []int, accept []int, params []Parameter) (map[string]interface{}, error) {
	for i, p := range params {
		for j := i + 1; j < len(params); j++ {
			if p.kind() == params[j].kind() {
				return nil, errors.New("duplicate parameter")
			}
		}
	}
	m := make(map[string]interface{})
	foundRequired := 0
paramLoop:
	for _, p := range params {
		for _, r := range require {
			if p.kind() == r {
				foundRequired++
				m[p.key()] = p.value()
				continue paramLoop
			}
		}
		for _, a := range accept {
			if p.kind() == a {
				m[p.key()] = p.value()
				continue paramLoop
			}
		}
		return nil, errors.New("unaccepted parameter")
	}
	if foundRequired != len(require) {
		return nil, errors.New("missing required parameter")
	}
	return m, nil
}

type ParameterInt struct {
	k       string
	v       int
	kindInt int
}

func (p ParameterInt) key() string {
	return p.k
}

func (p ParameterInt) value() interface{} {
	return p.v
}

func (p ParameterInt) kind() int {
	return p.kindInt
}

type ParameterInt64 struct {
	k       string
	v       int64
	kindInt int
}

func (p ParameterInt64) key() string {
	return p.k
}

func (p ParameterInt64) value() interface{} {
	return p.v
}

func (p ParameterInt64) kind() int {
	return p.kindInt
}
