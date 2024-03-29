// Code generated by github.com/DoNotPayHQ/gqlgen-fork, DO NOT EDIT.

package out

import (
	"fmt"
	"io"
	"strconv"
)

// InterfaceWithDescription is an interface with a description
type InterfaceWithDescription interface {
	IsInterfaceWithDescription()
}

type MissingInterface interface {
	IsMissingInterface()
}

type MissingUnion interface {
	IsMissingUnion()
}

// UnionWithDescription is an union with a description
type UnionWithDescription interface {
	IsUnionWithDescription()
}

type MissingInput struct {
	Name *string      `json:"name"`
	Enum *MissingEnum `json:"enum"`
}

type MissingTypeNotNull struct {
	Name     string               `json:"name"`
	Enum     MissingEnum          `json:"enum"`
	Int      MissingInterface     `json:"int"`
	Existing *ExistingType        `json:"existing"`
	Missing2 *MissingTypeNullable `json:"missing2"`
}

func (MissingTypeNotNull) IsMissingInterface()  {}
func (MissingTypeNotNull) IsExistingInterface() {}
func (MissingTypeNotNull) IsMissingUnion()      {}
func (MissingTypeNotNull) IsExistingUnion()     {}

type MissingTypeNullable struct {
	Name     *string             `json:"name"`
	Enum     *MissingEnum        `json:"enum"`
	Int      MissingInterface    `json:"int"`
	Existing *ExistingType       `json:"existing"`
	Missing2 *MissingTypeNotNull `json:"missing2"`
}

func (MissingTypeNullable) IsMissingInterface()  {}
func (MissingTypeNullable) IsExistingInterface() {}
func (MissingTypeNullable) IsMissingUnion()      {}
func (MissingTypeNullable) IsExistingUnion()     {}

// TypeWithDescription is a type with a description
type TypeWithDescription struct {
	Name *string `json:"name"`
}

func (TypeWithDescription) IsUnionWithDescription() {}

// EnumWithDescription is an enum with a description
type EnumWithDescription string

const (
	EnumWithDescriptionCat EnumWithDescription = "CAT"
	EnumWithDescriptionDog EnumWithDescription = "DOG"
)

var AllEnumWithDescription = []EnumWithDescription{
	EnumWithDescriptionCat,
	EnumWithDescriptionDog,
}

func (e EnumWithDescription) IsValid() bool {
	switch e {
	case EnumWithDescriptionCat, EnumWithDescriptionDog:
		return true
	}
	return false
}

func (e EnumWithDescription) String() string {
	return string(e)
}

func (e *EnumWithDescription) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EnumWithDescription(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EnumWithDescription", str)
	}
	return nil
}

func (e EnumWithDescription) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MissingEnum string

const (
	MissingEnumHello   MissingEnum = "Hello"
	MissingEnumGoodbye MissingEnum = "Goodbye"
)

var AllMissingEnum = []MissingEnum{
	MissingEnumHello,
	MissingEnumGoodbye,
}

func (e MissingEnum) IsValid() bool {
	switch e {
	case MissingEnumHello, MissingEnumGoodbye:
		return true
	}
	return false
}

func (e MissingEnum) String() string {
	return string(e)
}

func (e *MissingEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MissingEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MissingEnum", str)
	}
	return nil
}

func (e MissingEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
