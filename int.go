package incremental

import (
	"sync"
)

type Int struct {
	increment int
	lock      sync.Mutex
}

// Next returns the following 
func (i *Int) Next() int {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.increment++
	return i.increment
}

func (i *Int) Last() int {
	return i.increment
}
