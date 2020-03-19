package optional

import (
	"github.com/ThisIsClark/go-optional"
	"testing"
)

func TestOptional_Get(t *testing.T) {
	intData := 1
	op := optional.Optionalof(intData)
	intResData, ok := op.Get().(int)
	if !ok {
		t.Error("Not a int")
	}
	if intResData != 1 {
		t.Error("want a int 1, got ", intResData)
	}
	stringData := "TestString"
	op = optional.Optionalof(stringData)
	stringResData, ok := op.Get().(string)
	if !ok {
		t.Error("Not a string")
	}
	if stringResData != "TestString" {
		t.Error("want a string \"TestString\", got ", stringResData)
	}
}

func TestOptional_IsPresent(t *testing.T) {
	intData := 1
	op := optional.Optionalof(intData)
	if op.IsPresent() != true {
		t.Error("op should be present")
	}
	op = optional.OptionalEmpty()
	if op.IsPresent() != false {
		t.Error("op should be not present")
	}
}

func TestOptional_Equals(t *testing.T) {
	intData := 1
	op := optional.Optionalof(intData)
	if !op.Equals(intData) {
		t.Error("op should be equal to intData")
	}
}
