package module_event

import (
	"errors"
	ants "github.com/panjf2000/ants/v2"
	"github.com/slclub/go-tips/logf"
	"runtime"
	"time"
)

/**
 * 结偶性质比较强的组件
 * 异步 事件
 */

// --------------------------------------------------
// Event server
// --------------------------------------------------

/**
 *
 */
type Event struct {
	option     *Option
	container  *Container // 事件定义的容器
	flow       *workFlow  // 执行流
	singleChan chan *EventOper
	exit       chan struct{}
	ants_pool  AsyncSubmiter //*ants.Pool 默认引入
}

/**
 * Create an Event monitor.
 * You can create more than one.
 * you need to initialized completed because of no initializing method was offered to you.
 */
func NewEvent(option *Option) *Event {
	event := &Event{
		option:     option,
		container:  newContainer(),
		flow:       newWorkFlow(),
		singleChan: make(chan *EventOper, 10),
		exit:       make(chan struct{}),
	}
	event.init()
	return event
}

// 只能被调用一次
// 可以在类似sync.Once 中处理
func (self *Event) init() {
	// 启动单独routine 完成任务
	if self.option.InOrder {
		go self.single()
	}

	if self.option.Submiter == nil {
		// 启用默认 携程池
		self.initAnts()
	} else {
		// 使用自定义携程池
		self.ants_pool = self.option.Submiter
	}

	// 启动时间轮训routine
	if self.option.TimeTickPeriod > 0 {
		go self.timeListen()
	}
}

// 注册事件
func (self *Event) Register(handle EventHandle, eids ...EventValue) {
	if handle == nil || len(eids) == 0 {
		return
	}
	for _, eid := range eids {
		self.RegisterWithEventUnit(&EventItem{
			EID:    eid,
			Handle: handle,
		})
	}
}

func (self *Event) RegisterWithEventUnit(e *EventItem) error {
	if e == nil {
		return errors.New(ERROR_EVENT_ITEM_USED)
	}

	if e.Handle == nil {
		return errors.New(ERROR_HANDLE_IS_NIL)
	}
	if e.EID == nil {
		return errors.New(ERROR_HANDLE_EID_NIL)
	}
	self.container.Add(e)
	return nil
}

// 提交事件，并未立刻执行
func (self *Event) Submit(op *EventOper) {
	self.flow.submit(op)
}

// 同样是提交事件
func (self *Event) Trigger(eid EventValue, args ...any) {
	oper := &EventOper{EID: eid, Args: args}
	self.Submit(oper)
}

func (self *Event) Emit() {
	if self.flow.Len() == 0 {
		return
	}
	self.flow.idle(false)
	self.flow.Range(func(oper *EventOper) error {
		self.dealOne(oper)
		return nil
	})
	self.flow.done()
	self.flow.idle(true)
}

func (self *Event) dealOne(op *EventOper) {
	switch self.option.InOrder {
	case true:
		self.singleChan <- op
	case false:
		self.mutilple(op)
	}
}

// 这种写法，当发生panic 时，会出错，我们并没有处理recover
// Please be careful becase of I did not dealt the recover of panic in this routine
// It has a litte safe. if you're still worried，please check the panci error be come from your handle function
// or you can use another option selection.
func (self *Event) single() {
	for {
		select {
		case op, ok := <-self.singleChan:
			if !ok {
				return
			}
			self.container.Range(op.EID.Value(), func(oper *EventItem) error {
				oper.Handle(op)
				return nil
			})
		case <-self.exit:
			return
		}
	}
}

func (self *Event) timeListen() {
	ticker := time.NewTicker(self.option.TimeTickPeriod)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			self.Emit()
		case <-self.exit:
			return
		}
	}
}

func (self *Event) mutilple(op *EventOper) {
	self.container.Range(op.EID.Value(), func(item *EventItem) error {
		self.ants_pool.Submit(func() {
			item.Handle(op)
		})
		return nil
	})
}

func (self *Event) initAnts() {
	if self.ants_pool == nil {
		var err error = errors.New("")
		self.ants_pool, err = ants.NewPool(runtime.NumCPU()*8, ants.WithLogger(logf.Log()))
		if err != nil {
			panic(any(err))
		}
	}
	return
}

// 需要在服务进程关闭时调用
func (self *Event) Close() {
	// 线程池 是否释放
	if self.ants_pool != nil {
		switch pthread := self.ants_pool.(type) {
		case *ants.Pool:
			pthread.Release()
		}
	}

	close(self.singleChan)
	close(self.exit)
}

// --------------------------------------------------
// event container
// --------------------------------------------------

type Container struct {
	items map[string][]*EventItem
	// lock  sync.Locker
}

func newContainer() *Container {
	return &Container{
		items: make(map[string][]*EventItem),
	}
}

func (self *Container) Add(c *EventItem) {
	id := c.EID
	idvalue := id.Value()
	if _, ok := self.items[idvalue]; !ok {
		self.items[idvalue] = make([]*EventItem, 0, 1)
	}
	self.items[idvalue] = append(self.items[idvalue], c)
}

func (self *Container) Range(itemKey string, f func(item *EventItem) error) {
	items, ok := self.items[itemKey]
	if !ok {
		return
	}
	for _, item := range items {
		err := f(item)
		if err != nil {
			break
		}
	}
}

// --------------------------------------------------
// event global functions
// --------------------------------------------------
