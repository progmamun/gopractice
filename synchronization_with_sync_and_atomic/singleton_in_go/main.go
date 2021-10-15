package main

import "sync"

type obj struct{}

var instance *obj

func Get() *obj {
	if instance == nil {
		instance = &obj{}
	}
	return instance
}

//type obj struct{}

var (
	instance *obj
	lock     sync.Mutex
)

func Get() *obj {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &obj{}
	}
	return instance
}

type obj struct{}

var (
	instance *obj
	once     sync.Once
)

func Get() *obj {
	once.Do(func() {
		instance = &obj{}
	})
	return instance
}
