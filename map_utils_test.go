package map_utils

import (
	"math/cmplx"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {

	t.Run("invalid dict", func(t *testing.T) {
		var keys []string
		var dict = []string{"foo"}
		err := Keys(dict, &keys)
		assert.Error(t, err)
		assert.Empty(t, keys)
	})

	t.Run("empty map", func(t *testing.T) {
		var keys []string
		var dict = map[string]int{}
		err := Keys(dict, &keys)
		assert.NoError(t, err)
		assert.Empty(t, keys)
	})

	t.Run("slice value", func(t *testing.T) {
		var keys []string
		var dict = map[string]int{"foo": 123}
		err := Keys(dict, keys)
		assert.Error(t, err)
		assert.Empty(t, keys)
	})

	t.Run("invalid keys", func(t *testing.T) {
		var keys []string
		var dict = map[string]int{"foo": 123}
		err := Keys(dict, "invalid")
		assert.Error(t, err)
		assert.Empty(t, keys)
	})

	t.Run("type mismatch", func(t *testing.T) {
		var keys []int
		var dict = map[string]int{"foo": 123}
		err := Keys(dict, &keys)
		assert.Error(t, err)
		assert.Empty(t, keys)
	})

	t.Run("success", func(t *testing.T) {
		var keys []string
		var dict = map[string]int{"foo": 123, "baz": 456}
		err := Keys(dict, &keys)
		assert.NoError(t, err)
		assert.ElementsMatch(t, []string{"foo", "baz"}, keys)
	})
}

func TestValues(t *testing.T) {

	t.Run("invalid dict", func(t *testing.T) {
		var values []int
		var dict = []string{"foo"}
		err := Values(dict, &values)
		assert.Error(t, err)
		assert.Empty(t, values)
	})

	t.Run("empty map", func(t *testing.T) {
		var values []int
		var dict = map[string]string{}
		err := Values(dict, &values)
		assert.NoError(t, err)
		assert.Empty(t, values)
	})

	t.Run("slice value", func(t *testing.T) {
		var values []int
		var dict = map[string]int{"foo": 123}
		err := Values(dict, values)
		assert.Error(t, err)
		assert.Empty(t, values)
	})

	t.Run("invalid values", func(t *testing.T) {
		var values []int
		var dict = map[string]int{"foo": 123}
		err := Values(dict, 123)
		assert.Error(t, err)
		assert.Empty(t, values)
	})

	t.Run("type mismatch", func(t *testing.T) {
		var values []string
		var dict = map[string]int{"foo": 123}
		err := Values(dict, &values)
		assert.Error(t, err)
		assert.Empty(t, values)
	})

	t.Run("success", func(t *testing.T) {
		var values []int
		var dict = map[string]int{"foo": 123, "baz": 456}
		err := Values(dict, &values)
		assert.NoError(t, err)
		assert.ElementsMatch(t, []int{123, 456}, values)
	})
}

func TestSortedKeys(t *testing.T) {
	t.Run("invalid dict", func(t *testing.T) {
		var keys []string
		var dict = []string{"foo"}
		err := SortedKeys(dict, &keys)
		assert.Error(t, err)
		assert.Empty(t, keys)
	})

	t.Run("empty map", func(t *testing.T) {
		var keys []string
		var dict = map[string]int{}
		err := SortedKeys(dict, &keys)
		assert.NoError(t, err)
		assert.Empty(t, keys)
	})

	t.Run("slice value", func(t *testing.T) {
		var keys []string
		var dict = map[string]int{"foo": 123}
		err := SortedKeys(dict, keys)
		assert.Error(t, err)
		assert.Empty(t, keys)
	})

	t.Run("invalid keys", func(t *testing.T) {
		var keys []string
		var dict = map[string]int{"foo": 123}
		err := SortedKeys(dict, "invalid")
		assert.Error(t, err)
		assert.Empty(t, keys)
	})

	t.Run("type mismatch", func(t *testing.T) {
		var keys []int
		var dict = map[string]int{"foo": 123}
		err := SortedKeys(dict, &keys)
		assert.Error(t, err)
		assert.Empty(t, keys)
	})

	t.Run("success string", func(t *testing.T) {
		var keys []string
		var dict = map[string]int{"foo": 123, "baz": 456}
		err := SortedKeys(dict, &keys)
		assert.NoError(t, err)
		assert.Equal(t, []string{"baz", "foo"}, keys)
	})

	t.Run("success int", func(t *testing.T) {
		var keys []int
		var dict = map[int]string{2: "foo", 1: "bar"}
		err := SortedKeys(dict, &keys)
		assert.NoError(t, err)
		assert.Equal(t, []int{1, 2}, keys)
	})

	t.Run("success uint", func(t *testing.T) {
		var keys []uint
		var dict = map[uint]string{2: "foo", 1: "bar"}
		err := SortedKeys(dict, &keys)
		assert.NoError(t, err)
		assert.Equal(t, []uint{1, 2}, keys)
	})

	t.Run("success float", func(t *testing.T) {
		var keys []float64
		var dict = map[float64]string{2.0: "foo", 1.5: "bar"}
		err := SortedKeys(dict, &keys)
		assert.NoError(t, err)
		assert.Equal(t, []float64{1.5, 2.0}, keys)
	})

	t.Run("panic bool", func(t *testing.T) {
		var keys []bool
		var dict = map[bool]string{true: "foo", false: "bar"}
		assert.Panics(t, func() { SortedKeys(dict, &keys) })
	})

	t.Run("panic complex", func(t *testing.T) {
		var keys []complex128
		var dict = map[complex128]string{cmplx.Sqrt(-1 + 1i): "foo", 2i: "bar"}
		assert.Panics(t, func() { SortedKeys(dict, &keys) })
	})
}
