package incremental

import (
	"sync"
)

type Uint struct {
	increment uint
	lock      sync.Mutex
}

// Next returns with an integer that is exactly one higher as the previous call to Next() for this Uint
func (i *Uint) Next() uint {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.increment++
	return i.increment
}

// Last returns the number (uint) that was returned by the most recent call to this instance's Next()
func (i *Uint) Last() uint {
	return i.increment
}
