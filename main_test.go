package goconfig

import (
    "testing"
    "fmt"
)

func init() {
	//flag.BoolVar(&isTest, "t", false, "Run as Test")
}

func TestMain(t *testing.T) {
    if err := InitConficOnce("data.json", "sample.json", "test.json"); err != nil {
        fmt.Println(err)
        return
    }
  
    fmt.Println(GetString("Hallo"))
    fmt.Println(Get("Alter3"))
    fmt.Println(Get("Du", "Bist"))
    fmt.Println(Get("Du", "Gro"))
    fmt.Println(Get("Alter"))
    fmt.Println(GetString("Array2", "0", "a"))
    fmt.Println(Get("Array2", "1", "a"))
    fmt.Println(GetArrayInt("Array"))
}
