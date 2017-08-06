package interpreter

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestVarScope(t *testing.T) {
	asserts := assert.New(t)
	s := NewVarScope()
	var (
		err error
		val string
	)
	t.Run("Create", func(t *testing.T) {
		err = s.Create("a", "b")
		asserts.NoError(err)
		val, err = s.Get("a")
		asserts.NoError(err)
		asserts.Equal("b", val)
	})
	t.Run("Create error", func(t *testing.T) {
		err = s.Create("a", "b")
		asserts.Equal(ErrorVariableExists, err)
	})
	t.Run("Update", func(t *testing.T) {
		err = s.Update("a", "c")
		asserts.NoError(err)
		val, err = s.Get("a")
		asserts.NoError(err)
		asserts.Equal("c", val)
	})
	t.Run("Delete and Get error", func(t *testing.T) {
		err = s.Delete("a")
		asserts.NoError(err)
		val, err = s.Get("a")
		asserts.Equal(ErrorVariableDoesNotExist, err)
		asserts.Equal("undefined", val)
	})
	t.Run("Update error", func(t *testing.T) {
		err = s.Update("a", "c")
		asserts.Equal(ErrorVariableDoesNotExist, err)
	})
	t.Run("Delete error", func(t *testing.T) {
		err = s.Delete("a")
		asserts.Equal(ErrorVariableDoesNotExist, err)
	})
	t.Run("Set", func(t *testing.T) {
		s.Set("a", "z")
		val, err = s.Get("a")
		asserts.NoError(err)
		asserts.Equal("z", val)
	})
}