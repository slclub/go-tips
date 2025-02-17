package licenses

import (
	"errors"
	"time"
)

var GlobalTimeQueue *TimeQueue

// 防修改系统时间
type TimeQueue struct {
	simple        int64
	tick_duration time.Duration
	valid         bool
}

func NewTimeQueue() *TimeQueue {
	obj := &TimeQueue{
		simple:        time.Now().Unix(),
		tick_duration: time.Minute * 30,
		valid:         true,
	}
	return obj
}

func (self *TimeQueue) Go() {
	// 至少一分钟间隔
	if self.tick_duration < time.Minute {
		self.tick_duration = time.Minute
	}
	tick := time.NewTicker(self.tick_duration)
	defer tick.Stop()
	go func() {
		for {
			<-tick.C
			self.Tick()
			if !self.Valid() {
				panic(errors.New("the server  time was changed!!!"))
			}
		}
	}()
}

func (self *TimeQueue) Tick() {
	if self.simple == 0 {
		self.simple = time.Now().Unix()
		return
	}
	// server system time was changed.  so valid = false
	if self.simple > time.Now().Unix() {
		self.valid = false
		return
	}
	self.simple = time.Now().Unix()
}

func (self *TimeQueue) Valid() bool {
	return self.valid
}

// functions for time

func PingChangeTime() bool {
	GlobalTimeQueue.Tick()
	return GlobalTimeQueue.Valid()
}
