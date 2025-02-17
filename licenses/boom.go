package licenses

import (
	"errors"
	"fmt"
	"net"
	"os"
	"time"
)

type Boom struct {
	dur      time.Duration
	handles  []func() bool
	timeStop int // 停止次数
}

func NewBoom(dur time.Duration, fns ...func() bool) *Boom {
	if dur == 0 {
		dur = time.Hour
	}
	obj := &Boom{
		dur:     dur,
		handles: fns,
	}
	return obj
}

func (self *Boom) Run() {
	go self.update()
}

func (self *Boom) SetDur(d time.Duration) {
	self.dur = d
}

func (self *Boom) update() {
	d := time.NewTicker(self.dur)
	defer d.Stop()
	if !self.CheckInternet() {
		self.timeStop++
	}
	for {
		select {
		case <-d.C:
			if !self.CheckInternet() {
				self.timeStop++
			} else {
				self.timeStop = 0
			}
			if self.timeStop > 2 {
				err := errors.New(time.Now().Format("2006-01-02 15:04:05") + "License valide error. So the server was automatic stoped")
				fmt.Println(err)
				os.Exit(-1) // 退出进程
				//panic(err)
			}
		}
	}
}

func (self *Boom) CheckInternet() bool {
	for _, fn := range self.handles {
		if !fn() {
			return false
		}
	}
	return true
}

func PingBaidu() bool {
	return Ping("baidu.com")
}

func Ping(domain string) bool {
	timeout := 10 * time.Second // 设置超时时间

	succ := 0

	for i := 0; i < 10; i++ {
		conn, err := net.DialTimeout("tcp", domain+":80", timeout)
		if err != nil {
			fmt.Printf("Failed to connect the internet!!! \n")
			//fmt.Printf("Failed to connect to %s: %v\n", "intenet", err) // debug 测试使用
			continue
		}
		defer conn.Close()
		succ++
		break
	}

	return succ > 0
}
