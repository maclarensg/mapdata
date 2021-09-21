package mapdata

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type MapData map[string]interface{}
type ListData []interface{}

// Return a *MapData from a valid map[string]interface{}
// Which allows use to traverse and get data from a nested map
func NewMapData(data interface{}) (*MapData, error) {
	if reflect.TypeOf(data) != reflect.TypeOf(make(map[string]interface{})) {
		return nil, errors.New(fmt.Sprintf("Not a map[string]interface{}. %T", data))
	}
	rmap := MapData(data.(map[string]interface{}))
	return &rmap, nil
}

// Return the value by traversing the Map through the given path
func (d *MapData) GetPath(path string) (interface{}, error) {
	// We take a path `apple.promotion.sale.price` and convert to
	// an array of keys that helps to traverse the map ,
	// return the value if the path is valid
	return d.get(PathToArray(path))
}

// Recursive function for traversing the map
func (d *MapData) get(arr []string) (interface{}, error) {
	var value interface{}
	var ok bool

	key := arr[0] // Pop key, from array
	r := arr[1:]  // Remaining

	// Convert obj to map[string]interface{}
	// and check if the key exist
	// if not exist, return error
	dMap := map[string]interface{}(*d)
	value, ok = dMap[key]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Unable to find key %s", key))
	}

	// if we already reach the end of array
	// value is what we want, return it
	if len(r) == 0 {
		return value, nil
	}

	// Check if there is next key to search and value is a Map
	// If not, then there is nothing to search, return error
	if reflect.TypeOf(value).Kind() != reflect.Map {
		return nil, errors.New(fmt.Sprintf("Unable traverse map for key %s.", r[0]))
	}

	// Wrapper value as *MapData and we recurse for the remaing keys in array
	wrapper := MapData(value.(map[string]interface{}))
	mapPtr := &wrapper
	return mapPtr.get(r)
}

// Helper function that convert a path string to a string of keys
func PathToArray(path string) []string {
	return strings.Split(path, ".")
}
