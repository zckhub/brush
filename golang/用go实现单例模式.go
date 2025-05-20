package main

import "sync"

type Singleton struct {
}

var instance *Singleton
var lock sync.Mutex

func GetInstance() *Singleton {
	if instance == nil {
		lock.Lock()
		//instance = &Singleton{}
		defer lock.Unlock()
		if instance == nil {
			instance = &Singleton{}
		}
	}

	return instance
}
