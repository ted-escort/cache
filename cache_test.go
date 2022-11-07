package cache

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	str := "test"
	_, err := Set("key", []byte(str))
	if err != nil {
		fmt.Printf("error:=%#v\n", err)
		return
	}
}

func TestGet(t *testing.T) {
	cache, err := Get("key")
	if err != nil {
		fmt.Printf("error:=%#v\n", err)
		return
	}
	fmt.Printf("cache:=%#v\n", cache)
}

func TestDelete(t *testing.T) {
	err := Delete("key")
	if err != nil {
		fmt.Printf("error:%#v\n", err)
		return
	}
}
