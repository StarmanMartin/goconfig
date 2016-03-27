package goconfig

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

var (
	floatReg *regexp.Regexp
)

func init() {
	floatReg = regexp.MustCompile(`^\d*\.\d+$`)
}

type dataTypesEnum int
type keyMap map[string]key

const (
	intType dataTypesEnum = iota
	floatType
	stringType
	boolType
	configType
)

type config struct {
	intVal    []int
	floatVal  []float64
	stringVal []string
	boolVal   []bool
	subConfs  []*config
	keys      keyMap
}

type key struct {
	dataType         dataTypesEnum
	idx, rangeLength int
	isRange          bool
}

func (k keyMap) addKeyIfNotExists(mapKey string, dataType dataTypesEnum, idx int, isArray bool) {
	if keyElm, ok := k[mapKey]; !ok {
		k[mapKey] = key{dataType, idx, 1, isArray}
	} else {
		keyElm.rangeLength++
		k[mapKey] = keyElm
	}
}

func newConfig(mapData map[string]interface{}) *config {
	return newSubConfig(mapData)
}

func newSubConfig(mapData map[string]interface{}) *config {
	conf := &config{make([]int, 0), make([]float64, 0), make([]string, 0), make([]bool, 0), make([]*config, 0), make(map[string]key)}

	for i, volAsArray := range mapData {
		var (
			asArray []interface{}
			mapKey  = i
			typeVal = reflect.ValueOf(volAsArray) 
			isArray bool
		)

		if typeVal.Kind() == reflect.Slice {
			asArray = typeVal.Interface().([]interface{})
			isArray = true
		} else {
			asArray = []interface{}{volAsArray}
			isArray = false
		}

		for _, v := range asArray {
			typeVal := reflect.ValueOf(v)

			switch typeVal.Kind() {
			case reflect.Int, reflect.Float64, reflect.Float32:
				if floatReg.MatchString(fmt.Sprint(v)) {
					conf.keys.addKeyIfNotExists(mapKey, floatType, len(conf.floatVal), isArray)
					conf.floatVal = append(conf.floatVal, float64(typeVal.Float()))
				} else {
					conf.keys.addKeyIfNotExists(mapKey, intType, len(conf.intVal), isArray)
					conf.intVal = append(conf.intVal, int(typeVal.Float()))
				}
			case reflect.String:
				conf.keys.addKeyIfNotExists(mapKey, stringType, len(conf.stringVal), isArray)
				conf.stringVal = append(conf.stringVal, typeVal.String())
			case reflect.Bool:
				conf.keys.addKeyIfNotExists(mapKey, boolType, len(conf.boolVal), isArray)
				conf.boolVal = append(conf.boolVal, typeVal.Bool())
			case reflect.Map:
				subConf := typeVal.Interface().(map[string]interface{})
				conf.keys.addKeyIfNotExists(mapKey, configType, len(conf.subConfs), isArray)
				conf.subConfs = append(conf.subConfs, newSubConfig(subConf))
			}

		}
	}

	return conf
}

func (c *config) getIndexAndConfig(keyWords []string) (*config, key, bool) {
	temp := c
	var (
		tempKey key
		ok      bool
	)

	for i := 0; i < len(keyWords); i++ {
		val := keyWords[i]
		tempKey, ok = temp.keys[val]
		if !ok {
			return nil, tempKey, false
		}
		if tempKey.dataType == configType {
			if tempKey.isRange && i+1 < len(keyWords) {
				i++
				value, convErr := strconv.Atoi(keyWords[i])
				if convErr != nil || value >= tempKey.rangeLength {
					return nil, tempKey, false
				}

				tempKey.isRange = false
				temp = temp.subConfs[tempKey.idx+value]
			} else if tempKey.isRange {
				return temp, tempKey, true
			} else {
				temp = temp.subConfs[tempKey.idx]
			}
		} else if tempKey.isRange && i+2 == len(keyWords) {
			value, convErr := strconv.Atoi(keyWords[i+1])
			if convErr != nil || value >= tempKey.rangeLength {
				return nil, tempKey, false
			}

			tempKey.idx += value
			tempKey.isRange = false
			return temp, tempKey, true
		}
	}

	return temp, tempKey, true
}

func (c *config) get(keyWords ...string) (interface{}, bool) {
	temp, tempKey, ok := c.getIndexAndConfig(keyWords)

	if !ok || tempKey.dataType == configType{
		return nil, false
	}

	if tempKey.isRange {
		switch tempKey.dataType {
		case intType:
			return temp.intVal[tempKey.idx : tempKey.idx+tempKey.rangeLength], true
		case floatType:
			return temp.floatVal[tempKey.idx : tempKey.idx+tempKey.rangeLength], true
		case boolType:
			return temp.boolVal[tempKey.idx : tempKey.idx+tempKey.rangeLength], true
		case stringType:
			return temp.stringVal[tempKey.idx : tempKey.idx+tempKey.rangeLength], true
		case configType:
			return temp.subConfs[tempKey.idx : tempKey.idx+tempKey.rangeLength], true
		}
	} else {
		switch tempKey.dataType {
		case intType:
			return temp.intVal[tempKey.idx], true
		case floatType:
			return temp.floatVal[tempKey.idx], true
		case boolType:
			return temp.boolVal[tempKey.idx], true
		case stringType:
			return temp.stringVal[tempKey.idx], true
		case configType:
			return temp, true
		}
	}

	return nil, false
}

func (c *config) getString(keyWords ...string) (string, bool) {
	temp, tempKey, ok := c.getIndexAndConfig(keyWords)

	if !ok {
		return "", false
	}

	if tempKey.dataType == stringType {
		return temp.stringVal[tempKey.idx], true
	}

	return "", false
}

func (c *config) getFloat(keyWords ...string) (float64, bool) {
	temp, tempKey, ok := c.getIndexAndConfig(keyWords)

	if !ok {
		return 0, false
	}

	if tempKey.dataType == floatType {
		return temp.floatVal[tempKey.idx], true
	}

	return 0, false
}

func (c *config) getInt(keyWords ...string) (int, bool) {
	temp, tempKey, ok := c.getIndexAndConfig(keyWords)

	if !ok {
		return 0, false
	}

	if tempKey.dataType == intType {
		return temp.intVal[tempKey.idx], true
	}

	return 0, false
}

func (c *config) getSubConfig(keyWords ...string) (*config, bool) {
	temp, tempKey, ok := c.getIndexAndConfig(keyWords)

	if !ok {
		return nil, false
	}

	if tempKey.dataType == configType {
		return temp.subConfs[tempKey.idx], true
	}

	return nil, false
}

func (c *config) getBool(keyWords ...string) (bool, bool) {
	temp, tempKey, ok := c.getIndexAndConfig(keyWords)

	if !ok {
		return false, false
	}

	if tempKey.dataType == boolType {
		return temp.boolVal[tempKey.idx], true
	}

	return false, false
}

func (c *config) getArrayString(keyWords ...string) ([]string, bool) {
	temp, tempKey, ok := c.getIndexAndConfig(keyWords)

	if !ok {
		return nil, false
	}

	if tempKey.dataType == stringType {
		return temp.stringVal[tempKey.idx : tempKey.idx+tempKey.rangeLength], true
	}

	return nil, false
}

func (c *config) getArrayFloat(keyWords ...string) ([]float64, bool) {
	temp, tempKey, ok := c.getIndexAndConfig(keyWords)

	if !ok {
		return nil, false
	}

	if tempKey.dataType == floatType {
		return temp.floatVal[tempKey.idx : tempKey.idx+tempKey.rangeLength], true
	}

	return nil, false
}

func (c *config) getArrayInt(keyWords ...string) ([]int, bool) {
	temp, tempKey, ok := c.getIndexAndConfig(keyWords)

	if !ok {
		return nil, false
	}

	if tempKey.dataType == intType {
		return temp.intVal[tempKey.idx : tempKey.idx+tempKey.rangeLength], true
	}

	return nil, false
}

func (c *config) getArraySubConfig(keyWords ...string) ([]*config, bool) {
	temp, tempKey, ok := c.getIndexAndConfig(keyWords)

	if !ok {
		return nil, false
	}

	if tempKey.dataType == configType {
		return temp.subConfs[tempKey.idx : tempKey.idx+tempKey.rangeLength], true
	}

	return nil, false
}

func (c *config) getArrayBool(keyWords ...string) ([]bool, bool) {
	temp, tempKey, ok := c.getIndexAndConfig(keyWords)

	if !ok {
		return nil, false
	}

	if tempKey.dataType == boolType {
		return temp.boolVal[tempKey.idx : tempKey.idx+tempKey.rangeLength], true
	}

	return nil, false
}
