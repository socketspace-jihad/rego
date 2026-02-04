package core

import (
	"errors"
	"sync"
)

var (
	data sync.Map = sync.Map{}
)

func Upsert(key string, value any) {
	data.Store(key, value)
}

func Set(key string, value any) error {
	if _, ok := data.Load(key); ok {
		return errors.New("key already exists")
	}
	data.Store(key, value)
	return nil
}

func Get(key string) (any, error) {
	d, ok := data.Load(key)
	if !ok {
		return nil, errors.New("key doesn't exists")
	}
	return d, nil
}

func Delete(key string) {
	data.Delete(key)
}
