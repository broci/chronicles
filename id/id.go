package id

import "sync"

var mu sync.RWMutex

var n int64

func Next() int64 {
	incr()
	mu.RLock()
	v := n
	mu.RUnlock()
	return v
}

func incr() {
	mu.Lock()
	n++
	mu.Unlock()
}
