package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDB1(t *testing.T) {
	var val string
	var ok bool

	d := New()

	d.Set("ex", "10")
	val, _ = d.Get("ex")
	assert.Equal(t, "10", val)

	d.Unset("ex")
	_, ok = d.Get("ex")
	assert.False(t, ok)
}

func TestDB2(t *testing.T) {
	d := New()

	d.Set("a", "10")
	d.Set("b", "10")

	assert.Equal(t, 2, d.NumEqualTo("10"))

	d.Set("b", "30")

	assert.Equal(t, 1, d.NumEqualTo("10"))
}

func TestDBTransaction1(t *testing.T) {
	var val string
	var ok bool

	d := New()

	d.Begin()

	d.Set("a", "10")
	val, _ = d.Get("a")
	assert.Equal(t, "10", val)

	d.Begin()
	d.Set("a", "20")
	val, _ = d.Get("a")
	assert.Equal(t, "20", val)

	d.Rollback()
	val, _ = d.Get("a")
	assert.Equal(t, "10", val)

	d.Rollback()
	_, ok = d.Get("a")
	assert.False(t, ok)
}

func TestDBTransaction2(t *testing.T) {
	var val string

	d := New()
	d.Begin()
	d.Set("a", "30")
	d.Begin()
	d.Set("a", "40")
	d.Commit()

	val, _ = d.Get("a")
	assert.Equal(t, "40", val)

	assert.NotNil(t, d.Rollback())
}

func TestDBTransaction3(t *testing.T) {
	var val string
	var ok bool

	d := New()
	d.Set("a", "50")
	d.Begin()

	val, _ = d.Get("a")
	assert.Equal(t, "50", val)

	d.Set("a", "60")
	d.Begin()
	d.Unset("a")

	_, ok = d.Get("a")
	assert.False(t, ok)

	d.Rollback()
	val, ok = d.Get("a")
	assert.True(t, ok)
	assert.Equal(t, "60", val)

	d.Commit()
	val, _ = d.Get("a")
	assert.Equal(t, "60", val)
}

func TestDBTransaction4(t *testing.T) {
	d := New()

	d.Set("a", "10")
	d.Begin()
	assert.Equal(t, 1, d.NumEqualTo("10"))

	d.Begin()
	d.Unset("a")
	assert.Equal(t, 0, d.NumEqualTo("10"))

	d.Rollback()
	assert.Equal(t, 1, d.NumEqualTo("10"))

	d.Commit()
}
