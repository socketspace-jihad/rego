package core

import (
	"errors"
	"log"
	"testing"
)

var (
	IncorrectReturnValueErr = errors.New("storage doesn't show the correct value")
)

func Test_Storage(t *testing.T) {
	t.Run("test for store integer and then get", func(t *testing.T) {
		Set("test_key", 1)
		val, err := Get("test_key")
		if err != nil {
			t.Error(err)
		}
		if val != 1 {
			t.Error(IncorrectReturnValueErr)
		}
		Delete("test_key")
	})
	t.Run("test for store list of integer and then get", func(t *testing.T) {
		Set("test_key", []int{1, 2, 3, 4})
		val, err := Get("test_key")
		if err != nil {
			t.Error(err)
		}
		arr, ok := val.([]int)
		log.Println(arr)
		if !ok {
			t.Error(errors.New("cannot cast from any to list of integers"))
			return
		}
		if arr[0] != 1 {
			t.Error(IncorrectReturnValueErr)
		}
	})
	t.Run("test upsert twice", func(t *testing.T) {
		Upsert("test_key", 1)
		Upsert("test_key", "this is string")
		Delete("test_key")
	})
	t.Run("test for storing string and then get it", func(t *testing.T) {
		Set("test_key", "this is string")
		val, err := Get("test_key")
		if err != nil {
			t.Error(err)
		}
		if val != "this is string" {
			t.Error(IncorrectReturnValueErr)
		}
		Delete("test_key")
	})
	t.Run("test for storing map and then get it", func(t *testing.T) {
		Set("test_key", map[string]string{
			"a": "b",
		})
		val, err := Get("test_key")
		if err != nil {
			t.Error(err)
		}
		data, ok := val.(map[string]string)
		if !ok {
			t.Error(errors.New("cannot cast from any to map"))
			return
		}
		if data["a"] != "b" {
			t.Error(IncorrectReturnValueErr)
		}
	})
}
