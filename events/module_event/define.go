package module_event

import "strconv"

// --------------------------------------------------------
//  interface define
// --------------------------------------------------------

// event id interface
type EventValue interface {
	Value() string
}

// routine pool interface
// please remenber use Release method to free your memory of pool
type AsyncSubmiter interface {
	Submit(func()) error
	Release()
}

// --------------------------------------------------------
// type of event id  define
// --------------------------------------------------------

// 事件ID 分两种，都继承了 EventValue
type EVENT_ID_INT int
type EVENT_ID_STRING string

// 事件handle 函数
type EventHandle func(oper *EventOper)

func (self *EVENT_ID_INT) Value() string {
	return strconv.FormatInt(int64(*self), 10)
}

func (self *EVENT_ID_STRING) Value() string {
	return string(*self)
}

var _ EventValue = new(EVENT_ID_INT)
var _ EventValue = new(EVENT_ID_STRING)

// --------------------------------------------------------
// event id interface
// --------------------------------------------------------

const (
	ERROR_HANDLE_IS_NIL   = "[error] event handle is nil"
	ERROR_HANDLE_EID_NIL  = "[error] event handle is nil"
	ERROR_EVENT_ITEM_USED = "[error] event item is nil"
)
