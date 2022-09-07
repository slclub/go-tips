package safe

/**
 * @Email	:	slclub@163.com
 * @Author	:	xuyajun
 * @Description	: 	routine safe map
 */

import (
	"github.com/slclub/go-tips/spinlock"
	"sync"
)

/** 利用自旋锁实现的安全的map
 *  Secure map using spin lock
 *  routine security
 */
type QuickIntMap struct {
	data   map[int]any
	splock sync.Locker
}

func NewSafeMap() *QuickIntMap {
	return &QuickIntMap{
		data:   make(map[int]any),
		splock: spinlock.New(),
	}
}
func (this *QuickIntMap) Load(key int) (any, bool) {
	this.splock.Lock()
	defer this.splock.Unlock()
	val, ok := this.data[key]
	return val, ok
}

func (this *QuickIntMap) Store(key int, val any) {
	this.splock.Lock()
	defer this.splock.Unlock()
	this.data[key] = val
}

func (this *QuickIntMap) Range(fn func(key int, val any) bool) {
	for k, v := range this.data {
		rtn := fn(k, v)
		if rtn == false {
			break
		}
	}
}
