package spinlock

import (
	"runtime"
	"sync"
	"sync/atomic"
)

// =================================================================
// 自旋锁
// 借鉴了很多人的，源作者是谁暂时不清楚了，如有知道的请帮忙提issure
// I don't know who the original author is
// =================================================================

type spinLock uint32

func (sl *spinLock) Lock() {
	backoff := 1
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) {
		for i := 0; i < backoff; i++ {
			runtime.Gosched()
		}
		backoff <<= 1
	}
}
func (sl *spinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}

// NewSpinLock 实例化自旋锁
func New() sync.Locker {
	return new(spinLock)
}
