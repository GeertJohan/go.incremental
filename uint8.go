package incremental

import (
	"sync"
)

type Uint8 struct {
	increment uint8
	lock      sync.Mutex
}

// Next returns the following 
func (i *Uint8) Next() uint8 {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.increment++
	return i.increment
}

func (i *Uint8) Last() uint8 {
	return i.increment
}
