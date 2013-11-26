package incremental

import (
	"sync"
)

type Uint16 struct {
	increment uint16
	lock      sync.Mutex
}

// Next returns the following 
func (i *Uint16) Next() uint16 {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.increment++
	return i.increment
}

func (i *Uint16) Last() uint16 {
	return i.increment
}
