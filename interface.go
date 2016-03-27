package goconfig

import (
	"errors"
    "fmt"
)

var c *config

func InitConficOnce(mainConfig string, files ...string) error {
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
