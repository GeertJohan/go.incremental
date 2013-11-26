package incremental

import (
	"sync"
)

type Uint64 struct {
	increment uint64
	lock      sync.Mutex
}

// Next returns the following 
func (i *Uint64) Next() uint64 {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.increment++
	return i.increment
}

func (i *Uint64) Last() uint64 {
	return i.increment
}
