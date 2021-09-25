package mapdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockTestMap = map[string]interface{}{
	"a": map[string]interface{}{
		"a1": "alfa",
	},
	"b": []interface{}{
		map[string]interface{}{
			"b1": "beta1",
		},
		map[string]interface{}{
			"b2": "beta2",
		},
		map[string]interface{}{
			"b3": "beta3",
		},
		map[string]interface{}{
			"b4": "beta4",
		},
	},
	"c": []interface{}{
		"c1",
		"c2",
		"c3",
		"c4",
	},
	"d": "delta",
	"e": 5,
}

func TestNewMapData(t *testing.T) {
	assert := assert.New(t)
	mapTest := map[string]interface{}{
		"a": "alfa",
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
	assert := assert.New(t)
	var d interface{}
	var err error

	res, err := NewMapData(mockTestMap)
	assert.NotNil(res)
	assert.NoError(err)

	d, err = res.get([]string{"a", "a1"})
	assert.NoError(err)
	assert.Equal("alfa", d)

	d, err = res.get([]string{"a", "b"})
	assert.Error(err)
	assert.Nil(d)

	d, err = res.get([]string{"b", "c", "d"})
	assert.Error(err)
	assert.Nil(d)
}

func TestGetPathFromMapData(t *testing.T) {

	assert := assert.New(t)
	var d interface{}
	var err error

	res, err := NewMapData(mockTestMap)
	assert.NotNil(res)
	assert.NoError(err)

	d, err = res.GetPath("a.a1")
	assert.NoError(err)
	assert.Equal("alfa", d)

	d, err = res.GetPath("a.b")
	assert.Error(err)
	assert.Nil(d)
}

func TestGetPathValueString(t *testing.T) {
	assert := assert.New(t)
	res, err := NewMapData(mockTestMap)
	assert.NoError(err)
	val1, err := res.GetPathValueString("d")
	assert.NoError(err)
	assert.Equal(mockTestMap["d"], val1)
	val2, err := res.GetPathValueString("e")
	assert.Error(err)
	assert.Equal("", val2)
	val3, err := res.GetPathValueString("a")
	assert.Error(err)
	assert.Equal("", val3)
	val4, err := res.GetPathValueString("b")
	assert.Error(err)
	assert.Equal("", val4)
	val5, err := res.GetPathValueString("c")
	assert.Error(err)
	assert.Equal("", val5)
	val6, err := res.GetPathValueString("a.a1.a2")
	assert.Error(err)
	assert.Equal("", val6)
}

func TestGetPathValueMap(t *testing.T) {
	assert := assert.New(t)
	res, err := NewMapData(mockTestMap)
	assert.NoError(err)
	val1, err := res.GetPathValueMap("a")
	assert.NoError(err)
	assert.IsType(&MapData{}, val1)
	val2, err := res.GetPathValueMap("b")
	assert.Error(err)
	assert.Nil(val2)
	val3, err := res.GetPathValueMap("c")
	assert.Error(err)
	assert.Nil(val3)
	val4, err := res.GetPathValueMap("d")
	assert.Error(err)
	assert.Nil(val4)
	val5, err := res.GetPathValueMap("e")
	assert.Error(err)
	assert.Nil(val5)
	val6, err := res.GetPathValueMap("a.a1.a2")
	assert.Error(err)
	assert.Nil(val6)
}

func TestGetPathValueListMap(t *testing.T) {
	assert := assert.New(t)

	res, err := NewMapData(mockTestMap)
	assert.NoError(err)
	val1, err := res.GetPathValueListMap("b")
	assert.NoError(err)
	assert.Equal(len(mockTestMap["b"].([]interface{})), len(val1))
	val2, err := res.GetPathValueListMap("a")
	assert.Error(err)
	assert.Empty(val2)
	val3, err := res.GetPathValueListMap("c")
	assert.Error(err)
	assert.Empty(val3)
	val4, err := res.GetPathValueListMap("d")
	assert.Error(err)
	assert.Empty(val4)
	val5, err := res.GetPathValueListMap("e")
	assert.Error(err)
	assert.Empty(val5)
	val6, err := res.GetPathValueListMap("b.c.d")
	assert.Error(err)
	assert.Empty(val6)
}
