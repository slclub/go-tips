package module_event

import "time"

// --------------------------------------------------
// EventItem 完整的事件单元
// --------------------------------------------------

type EventItem struct {
	EID    EventValue
	Handle EventHandle
	Desc   string // 备注描述
}

// --------------------------------------------------
// EventOper is like a courier to translate paramters
// --------------------------------------------------

type EventOper struct {
	EID  EventValue
	Oper string //间接 体提供给 事件函数特殊标示，每个事件ID 可以在同一次被触发多次
	Args []any
}

// --------------------------------------------------
// Event Options
// --------------------------------------------------

/**
 * The option struct is very important. It has a great influence on Event Monitor.
 */
type Option struct {
	// 按顺序 true:单线程，false:并发多线程
	InOrder bool

	// 按时间监听 ， 定时发射事件，执行Emit
	// 值 <= 0 : 不启动按时间监听；此种情况需要您自己 把握启动事件的时机
	// 值 > 0 : 按TimeTickPeriod 纳秒 轮训一次 且执行Emit 方法
	TimeTickPeriod time.Duration //

	// 并发异步提交，线程池管理
	// 如果想用自己的线程池可以 嵌入进来
	Submiter AsyncSubmiter
}

// --------------------------------------------------
// some  functions that are using frequently
// --------------------------------------------------

func HandleConvert(fn func(args ...any)) EventHandle {
	return func(e *EventOper) {
		fn(e.Args...)
	}
}
