package main

import "sync"

type keyQueue struct {
	keys    []string
	maxSize int
	sync.Mutex
}

func newKeyQueue(size int) *keyQueue {
	return &keyQueue{
		keys:    []string{},
		maxSize: size,
	}
}

func (k *keyQueue) add(key string) {
	k.Lock()
	defer k.Unlock()

	if len(k.keys) == k.maxSize {
		k.keys = k.keys[1:]
	}

	k.keys = append(k.keys, key)
}

func (k *keyQueue) doesExist(key string) bool {
	k.Lock()
	defer k.Unlock()

	for _, queueKey := range k.keys {
		if key == queueKey {
			return true
		}
	}

	return false
}
