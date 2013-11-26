package incremental

import (
	"sync"
)

type Uint32 struct {
	increment uint32
	lock      sync.Mutex
}

// Next returns the following 
func (i *Uint32) Next() uint32 {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.increment++
	return i.increment
}

func (i *Uint32) Last() uint32 {
	return i.increment
}
