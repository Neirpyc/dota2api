package dota2api

import (
	"errors"
	"fmt"
)

type parameterKind int

const (
	parameterKindHeroId parameterKind = iota
	parameterKindMatchesRequested
	parameterKindAccountId
	parameterKindStartAtMatchId
	parameterKindMinPlayers
	parameterKindMatchId
	parameterStartMatchAtSeqNum
	parameterSteamIds
	parameterSteamId
	parameterVanityUrl
)

func (p parameterKind) String() string {
	switch p {
	case parameterKindHeroId:
		return "hero Id"
	case parameterKindMatchesRequested:
		return "matches requested"
	case parameterKindAccountId:
		return "account Id"
	case parameterKindStartAtMatchId:
		return "start at match Id"
	case parameterKindMinPlayers:
		return "min players"
	case parameterKindMatchId:
		return "match Id"
	case parameterStartMatchAtSeqNum:
		return "start match at seq num"
	case parameterSteamIds:
		return "steam Ids"
	case parameterSteamId:
		return "steam Id"
	case parameterVanityUrl:
		return "vanity Url"
	}
	return "unknown"
}

type Parameter interface {
	key() string
	value() interface{}
	kind() parameterKind
}

func getParameterMap(require []parameterKind, accept []parameterKind, params []Parameter) (map[string]interface{}, error) {
	for i, p := range params {
		for j := i + 1; j < len(params); j++ {
			if p.kind() == params[j].kind() {
				return nil, errors.New(fmt.Sprintf("duplicate parameter \"%s\"", p.kind().String()))
			}
		}
	}
	m := make(map[string]interface{})
	reqState := make([]bool, len(require))
	reqFound := 0
	var unaccepted []string
paramLoop:
	for _, p := range params {
		for i, r := range require {
			if p.kind() == r {
				reqState[i] = true
				reqFound++
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
		unaccepted = append(unaccepted, p.kind().String())
	}
	if unaccepted != nil {
		return nil, errors.New(fmt.Sprintf("unaccepted parameter(s) %v", unaccepted))
	}
	if reqFound != len(require) {
		var missing []string
		for i, p := range reqState {
			if !p {
				missing = append(missing, require[i].String())
			}
		}
		return nil, errors.New(fmt.Sprintf("missing required parameter(s) %v", missing))
	}
	return m, nil
}

type ParameterInt struct {
	k       string
	v       int
	kindInt parameterKind
}

func (p ParameterInt) key() string {
	return p.k
}

func (p ParameterInt) value() interface{} {
	return p.v
}

func (p ParameterInt) kind() parameterKind {
	return p.kindInt
}

type ParameterString struct {
	k       string
	v       string
	kindInt parameterKind
}

func (p ParameterString) key() string {
	return p.k
}

func (p ParameterString) value() interface{} {
	return p.v
}

func (p ParameterString) kind() parameterKind {
	return p.kindInt
}

type ParameterInt64 struct {
	k       string
	v       int64
	kindInt parameterKind
}

func (p ParameterInt64) key() string {
	return p.k
}

func (p ParameterInt64) value() interface{} {
	return p.v
}

func (p ParameterInt64) kind() parameterKind {
	return p.kindInt
}
