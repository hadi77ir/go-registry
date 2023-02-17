package registry

import (
	"testing"
)

var testData = map[string]string{
	"name1": "value1",
	"name2": "value2",
	"name3": "value3",
	"name4": "value4",
}
var nonExistingKeys = []string{"name5", "name6", "name7"}

func getRegistryEmpty() *Registry[string] {
	return &Registry[string]{}
}

func getRegistryFilled() *Registry[string] {
	r := &Registry[string]{}
	for k, v := range testData {
		r.Register(k, v)
	}
	return r
}

func TestRegistry_Get(t *testing.T) {
	r := getRegistryFilled()
	// check for existing data
	for k, v := range testData {
		if value, found := r.Get(k); !found || value != v {
			t.Fail()
			return
		}
	}
	// check for non-existing data
	for _, k := range nonExistingKeys {
		if value, found := r.Get(k); found || value != "" {
			t.Fail()
			return
		}
	}
}

func TestRegistry_Keys(t *testing.T) {
	r := getRegistryFilled()
	keys := r.Keys()
	for _, k := range keys {
		if _, found := testData[k]; !found {
			t.Fail()
			return
		}
	}
	if len(r.items) != len(testData) {
		t.Fail()
		return
	}
}

func TestRegistry_KeysEmpty(t *testing.T) {
	r := getRegistryEmpty()
	keys := r.Keys()
	if len(keys) != 0 {
		t.Fail()
	}
}
