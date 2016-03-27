# goconfig

A small go config JSON tool

## What is goconfig

goconfig is another config tool for golang. reads a json file and provides the content converted.

## Install goconfig

```bash
go get github.com/starmanmaritn/goconfig
```

```go
import "github.com/starmanmaritn/goconfig"

func main() {
    goconfig.InitConficOnce("config.json", "sample.json")    
}
```

## Methods

* [`InitConficOnce`](#InitConficOnce)
* [`Get`](#Get)
* [`GetInt`](#GetInt)
* [`GetFloat`](#GetFloat)
* [`GetBool`](#GetBool)
* [`GetString`](#GetString)
* [`GetArrayInt`](#GetArrayInt)
* [`GetArrayFloat`](#GetArrayFloat)
* [`GetArrayBool`](#GetArrayBool)
* [`GetArrayString`](#GetArrayString)

<a name="InitConficOnce"></a>

### InitConficOnce(mainConfig string, files ...string) error

InitConficOnce initialises the configuration container. It reads all files starting with the mainConfig.
The most right parameter is the strongest which means that if a value has set before it gets overwritten.

#### Parameter

* `mainConfig` *string* Absolute path to config file
* `files` *string* List of absolute paths to sub config files

#### return

* `error` nil if success. Else if mainConfig is not readable

<a name="Get"></a>

### Get(keyWords ...string) (interface{}, bool)

Get returns a value of th config container. The method gets a JSON path split als string. For example:

<a name="SampleJson"></a>

JSON:

```json
{
    "sample": {
        "array": [
            {
              "name": "Joo"
            },
            {
              "name": "Too"
            }
        ]
    }
}
```

GO:

```go
name, err := goconfig.GetString("sample", "array", "1", "name")
fmt.Println("name:", name)
```

Result:

```bash
> name: Too
```

#### Parameter

* `keyWords` *...string* a JSON path to a value

#### return

* `interface{}` value of a field in the JSON file as interface{}
* `bool` true if path is a correct path to a value

<a name="GetInt"></a>

### GetInt(keyWords ...string) (int, bool)

GetInt works as the ['Get'](#Get) method but returns just integer.

#### Parameter

* `keyWords` *...string* a JSON path to a value

#### return

* `int` value of a field in the JSON file if value is an integer
* `bool` true if the path is a correct path to an integer value

<a name="GetFloat"></a>

### GetFloat(keyWords ...string) (float, bool)

GetFloat works as the ['Get'](#Get) method but returns just float.

#### Parameter

* `keyWords` *...string* a JSON path to a value

#### return

* `float` value of a field in the JSON file if value is a float
* `bool` true if the path is a correct path to a float value

<a name="GetBool"></a>

### GetBool(keyWords ...string) (bool, bool)

GetBool works as the ['Get'](#Get) method but returns just boolean.

#### Parameter

* `keyWords` *...string* a JSON path to a value

#### return

* `bool` value of a field in the JSON file if value is a boolean
* `bool` true if the path is a correct path to a boolean value

<a name="GetString"></a>

### GetString(keyWords ...string) (string, bool)

GetString works as the ['Get'](#Get) method but returns just string.

#### Parameter

* `keyWords` *...string* a JSON path to a value

#### return

* `string` value of a field in the JSON file if value is a string
* `bool` true if the path is a correct path to a string value

<a name="GetArrayInt"></a>

### GetArrayInt(keyWords ...string) ([]int, bool)

GetArrayInt works as the ['Get'](#Get) method but returns just integer slices.

#### Parameter

* `keyWords` *...string* a JSON path to a value

#### return

* `[]int` value of a field in the JSON file if value is an integer slice
* `bool` true if the path is a correct path to an integer slice

<a name="GetArrayFloat"></a>

### GetArrayFloat(keyWords ...string) ([]float, bool)

GetArrayFloat works as the ['Get'](#Get) method but returns just float slices.

#### Parameter

* `keyWords` *...string* a JSON path to a value

#### return

* `[]float` value of a field in the JSON file if value is a float slice
* `bool` true if the path is a correct path to a float slice

<a name="GetArrayBool"></a>

### GetArrayBool(keyWords ...string) ([]bool, bool)

GetArrayBool works as the ['Get'](#Get) method but returns just boolean slices.

#### Parameter

* `keyWords` *...string* a JSON path to a value

#### return

* `[]bool` value of a field in the JSON file if value is a boolean slice
* `bool` true if the path is a correct path to a boolean slice

<a name="GetArrayString"></a>

### GetArrayString(keyWords ...string) ([]string, bool)

GetArrayString works as the ['Get'](#Get) method but returns just string slices.

#### Parameter

* `keyWords` *...string* a JSON path to a value

#### return

* `[]string` value of a field in the JSON file if value is a string slice
* `bool` true if the path is a correct path to a string slice