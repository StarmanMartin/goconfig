package goconfig

import (
	"errors"
	"fmt"
	"log"
)

var c *config

func InitConficOnce(mainConfig string, files ...string) error {
	return InitConfigOnce(mainConfig, files...)
}

func InitConfigOnce(mainConfig string, files ...string) error {
	val, err := readJSON(mainConfig)
	if err != nil {
		return errors.New("Main config: " + err.Error())
	}

	for _, file := range files {

		if value, errI := readJSON(file); errI == nil {
			val = margeJSON(val, value)
		} else {
			fmt.Println(errI)
		}
	}

	c = newConfig(val)

	return nil
}

func Get(keyWords ...string) (interface{}, bool) {
	if c == nil {
		return nil, false
	}

	return c.get(keyWords...)
}

func GetString(keyWords ...string) (string, bool) {
	if c == nil {
		return "", false
	}

	return c.getString(keyWords...)
}

func GetFloat(keyWords ...string) (float64, bool) {
	if c == nil {
		return 0, false
	}

	return c.getFloat(keyWords...)
}

func GetInt(keyWords ...string) (int, bool) {
	if c == nil {
		return 0, false
	}

	return c.getInt(keyWords...)
}

func GetBool(keyWords ...string) (bool, bool) {
	if c == nil {
		return false, false
	}

	return c.getBool(keyWords...)
}

func GetArrayString(keyWords ...string) ([]string, bool) {
	if c == nil {
		return nil, false
	}

	return c.getArrayString(keyWords...)
}

func GetArrayFloat(keyWords ...string) ([]float64, bool) {
	if c == nil {
		return nil, false
	}

	return c.getArrayFloat(keyWords...)
}

func GetArrayInt(keyWords ...string) ([]int, bool) {
	if c == nil {
		return nil, false
	}

	return c.getArrayInt(keyWords...)
}

func GetArrayBool(keyWords ...string) ([]bool, bool) {
	if c == nil {
		return nil, false
	}

	return c.getArrayBool(keyWords...)
}

func MustGetString(keyWords ...string) (string) {
	if c == nil {
		return ""
	}

	val, ok := c.getString(keyWords...)
	if !ok {
		log.Panic("Value not found!")
	}

	return val
}

func MustGetFloat(keyWords ...string) (float64) {
	if c == nil {
		return 0
	}

	val, ok := c.getFloat(keyWords...)
	if !ok {
		log.Panic("Value not found!")
	}

	return val
}

func MustGetInt(keyWords ...string) (int) {
	if c == nil {
		return 0
	}

	val, ok := c.getInt(keyWords...)
	if !ok {
		log.Panic("Value not found!")
	}

	return val
}

func MustGetBool(keyWords ...string) (bool) {
	if c == nil {
		return false
	}

	val, ok := c.getBool(keyWords...)
	if !ok {
		log.Panic("Value not found!")
	}

	return val
}

func MustGetArrayString(keyWords ...string) ([]string) {
	if c == nil {
		return nil
	}

	val, ok := c.getArrayString(keyWords...)
	if !ok {
		log.Panic("Value not found!")
	}

	return val
}

func MustGetArrayFloat(keyWords ...string) ([]float64) {
	if c == nil {
		return nil
	}

	val, ok := c.getArrayFloat(keyWords...)
	if !ok {
		log.Panic("Value not found!")
	}

	return val
}

func MustGetArrayInt(keyWords ...string) ([]int) {
	if c == nil {
		return nil
	}

	val, ok := c.getArrayInt(keyWords...)
	if !ok {
		log.Panic("Value not found!")
	}

	return val
}

func MustGetArrayBool(keyWords ...string) ([]bool) {
	if c == nil {
		return nil
	}

	val, ok := c.getArrayBool(keyWords...)
	if !ok {
		log.Panic("Value not found!")
	}

	return val
}
