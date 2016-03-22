package goconfig

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
)

func readJSON(file string) (map[string]interface{}, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var slice map[string]interface{}
	err = json.Unmarshal(data, &slice)
	if err != nil {
		return nil, err
	}

	return slice, nil
}

func margeJSON(a, b map[string]interface{}) map[string]interface{} {
	for i, va := range a {
		if vb, ok := b[i]; ok {
			typeValA := reflect.ValueOf(va)
			typeValB := reflect.ValueOf(vb)

			if typeValA.Kind() == reflect.Map && typeValA.Kind() == typeValB.Kind() {
				mapA := typeValA.Interface().(map[string]interface{})
				mapB := typeValB.Interface().(map[string]interface{})
				vb = margeJSON(mapA, mapB)
			}

			a[i] = vb
		}
	}

	for i, vb := range b {
		if va, ok := a[i]; ok {
			typeValA := reflect.ValueOf(va)
			typeValB := reflect.ValueOf(vb)

			if typeValA.Kind() == reflect.Map && typeValA.Kind() == typeValB.Kind() {
				mapA := typeValA.Interface().(map[string]interface{})
				mapB := typeValB.Interface().(map[string]interface{})
				vb = margeJSON(mapA, mapB)
			}
		}

		a[i] = vb
	}

	return a
}
