package types

import (
	"fmt"
)

type ObjectType int

type typeProp struct {
	typ  ObjectType
	name string
}

const (
	TyNumber ObjectType = iota
	TyString
	TyClosure
	TyNil
	TySymbol
	TyPair
	TyBoolean
	TySyntax
	TyContinuation
	TyVector
	TyUndefined
	TyError

	TyCallInfo // for internal use
)

var typeProps = []*typeProp{
	&typeProp{TyNumber, "number"},
	&typeProp{TyString, "string"},
	&typeProp{TyClosure, "closure"},
	&typeProp{TyNil, "nil"},
	&typeProp{TySymbol, "symbol"},
	&typeProp{TyPair, "pair"},
	&typeProp{TyBoolean, "boolean"},
	&typeProp{TySyntax, "syntax"},
	&typeProp{TyContinuation, "continuation"},
	&typeProp{TyVector, "vector"},
	&typeProp{TyUndefined, "undefined"},
	&typeProp{TyError, "error"},
	&typeProp{TyCallInfo, "callinfo"},
}

type Object interface {
	String() string
	Type() ObjectType
}

type Slicable interface {
	Slice() ([]Object, error)
}

type SlicableObject interface {
	Object
	Slicable
}

type (
	Number float64
	String string
	Nil    struct{}
	Symbol struct {
		Name String
	}
	Boolean   bool
	Undefined struct{}
)

func (num Number) String() string {
	return fmt.Sprint(float64(num))
}

func (num Number) Type() ObjectType { return TyNumber }

func (s String) String() string {
	return string(s)
}

func (s String) Type() ObjectType {
	return TyString
}

func NewSymbol(name string) *Symbol {
	return &Symbol{Name: String(name)}
}

func (s *Symbol) String() string {
	return s.Name.String()
}

func (s *Symbol) Type() ObjectType {
	return TySymbol
}

func (n *Nil) String() string {
	return "()"
}

func (n *Nil) Type() ObjectType {
	return TyNil
}

func (n *Nil) Slice() ([]Object, error) {
	return []Object{}, nil
}

var NilObject = &Nil{}

func (b Boolean) Type() ObjectType {
	return TyBoolean
}

func (b Boolean) String() string {
	if b {
		return "#t"
	}
	return "#f"
}

func (un *Undefined) Type() ObjectType {
	return TyUndefined
}

func (un *Undefined) String() string {
	return "undefined"
}

var UndefinedObject = &Undefined{}
