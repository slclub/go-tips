package module_event

import (
	"github.com/slclub/go-tips/spinlock"
	"sync"
)

// -------------------------------------------------
// 执行流
// -------------------------------------------------
type workFlow struct {
	sequence     *sequenceBox // 执行序列
	idleRun      bool         // true:空闲, false: running
	sequenceBack *sequenceBox
	mu           sync.Locker
}

func newWorkFlow() *workFlow {
	return &workFlow{
		sequence:     newSequenceBox(),
		idleRun:      true,
		sequenceBack: newSequenceBox(),
		mu:           spinlock.New(),
	}
}

func (self *workFlow) idle(runs ...bool) bool {
	self.mu.Lock()
	defer self.mu.Unlock()

	if len(runs) == 0 {
		return self.idleRun
	}

	self.idleRun = runs[0]
	return self.idleRun
}

// 提交到工作流函数
func (self *workFlow) submit(op *EventOper) {
	self.mu.Lock()
	defer self.mu.Unlock()
	if self.idleRun {
		self.sequence.add(op)
		return
	}
	self.sequenceBack.add(op)
}

func (self *workFlow) done() {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.sequence = self.sequenceBack
	self.sequenceBack = newSequenceBox()
}

func (self *workFlow) Range(fn func(op *EventOper) error) {
	self.sequence.Range(fn)
}

func (self *workFlow) Len() int {
	return self.sequence.Len()
}

// -------------------------------------------------
// 执行流盒子
// -------------------------------------------------

type sequenceBox struct {
	opers []*EventOper
	lock  sync.Locker
}

func newSequenceBox() *sequenceBox {
	return &sequenceBox{
		lock:  spinlock.New(),
		opers: make([]*EventOper, 0, 5),
	}
}

func (self *sequenceBox) add(op *EventOper) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.opers = append(self.opers, op)
}

func (self *sequenceBox) Range(fn func(op *EventOper) error) {
	if fn == nil {
		return
	}
	for _, oper := range self.opers {
		err := fn(oper)
		if err != nil {
			break
		}
	}
}

func (self *sequenceBox) Len() int {
	return len(self.opers)
}
