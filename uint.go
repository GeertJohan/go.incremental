package incremental

import (
	"sync"
)

type Uint struct {
	increment uint
	lock      sync.Mutex
}

// Next returns the following 
func (i *Uint) Next() uint {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.increment++
	return i.increment
}

func (i *Uint) Last() uint {
	return i.increment
}
