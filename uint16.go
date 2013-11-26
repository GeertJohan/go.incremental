package incremental

import (
	"sync"
)

type Uint16 struct {
	increment uint16
	lock      sync.Mutex
}

// Next returns with an integer that is exactly one higher as the previous call to Next() for this Uint16
func (i *Uint16) Next() uint16 {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.increment++
	return i.increment
}

// Last returns the number (uint16) that was returned by the most recent call to this instance's Next()
func (i *Uint16) Last() uint16 {
	return i.increment
}
