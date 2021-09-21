package mapdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMapData(t *testing.T) {
	assert := assert.New(t)
	mapTest := map[string]interface{}{
		"a": "alpha",
		"b": "beta",
		"c": "charlie",
		"d": "delta",
	}

	res, err := NewMapData(mapTest)
	assert.NotNil(res)
	assert.NoError(err)
	res, err = NewMapData([]string{"a", "b", "c", "d"})
	assert.Error(err)
	assert.Nil(res)
}

func TestPathToArray(t *testing.T) {
	assert := assert.New(t)
	path := "a.b.c.d"
	expect := []string{"a", "b", "c", "d"}
	assert.Equal(expect, PathToArray(path))
}

func TestGetFromMapData(t *testing.T) {
	mapTest := map[string]interface{}{
		"a": map[string]interface{}{
			"a1": "alpha",
		},
		"b": "beta",
		"c": "charlie",
		"d": "delta",
	}
	assert := assert.New(t)
	var d interface{}
	var err error

	res, err := NewMapData(mapTest)
	assert.NotNil(res)
	assert.NoError(err)

	d, err = res.get([]string{"a", "a1"})
	assert.NoError(err)
	assert.Equal("alpha", d)

	d, err = res.get([]string{"a", "b"})
	assert.Error(err)
	assert.Nil(d)

	d, err = res.get([]string{"b", "c", "d"})
	assert.Error(err)
	assert.Nil(d)
}

func TestGetPathFromMapData(t *testing.T) {
	mapTest := map[string]interface{}{
		"a": map[string]interface{}{
			"a1": "alpha",
		},
		"b": "beta",
		"c": "charlie",
		"d": "delta",
	}
	assert := assert.New(t)
	var d interface{}
	var err error

	res, err := NewMapData(mapTest)
	assert.NotNil(res)
	assert.NoError(err)

	d, err = res.GetPath("a.a1")
	assert.NoError(err)
	assert.Equal("alpha", d)

	d, err = res.GetPath("a.b")
	assert.Error(err)
	assert.Nil(d)
}
