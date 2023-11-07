package module_event

import (
	"testing"
	"time"
)

func TestEventCommon(t *testing.T) {
	eventMonitor := NewEvent(&Option{})

	// register
	eventMonitor.Register(EventHandle(handleLogin), &TEST_EVENT_ID_LOGIN)                        // register fail
	eventMonitor.Register(EventHandle(handleLogout), &TEST_EVENT_ID_LOGOUT)                      // register fail
	eventMonitor.Register(EventHandle(handleTrace), &TEST_EVENT_ID_LOGIN, &TEST_EVENT_ID_LOGOUT) // regsiter fail
	eventMonitor.Register(EventHandle(handleLogout), nil)                                        // register fail
	eventMonitor.Register(HandleConvert(anotherHandle), &TEST_EVENT_ID_ANOTHER)                  // register ok

	// trigger
	eventMonitor.Submit(&EventOper{&TEST_EVENT_ID_LOGIN, "", []any{t, 1, 2, "event"}}) // ok
	eventMonitor.Trigger(&TEST_EVENT_ID_ANOTHER, t, 3, "gold")

	// emit
	eventMonitor.Emit()
	time.Sleep(10 * time.Millisecond)

	// release
	eventMonitor.Close()
}

func TestEventSingle(t *testing.T) {
	eventMonitor := NewEvent(&Option{InOrder: true})

	// register
	eventMonitor.Register(EventHandle(handleLogin), &TEST_EVENT_ID_LOGIN)   // register fail
	eventMonitor.Register(EventHandle(handleLogout), &TEST_EVENT_ID_LOGOUT) // register fail

	// trigger
	eventMonitor.Submit(&EventOper{&TEST_EVENT_ID_LOGIN, "", []any{t, 1, 2, "event"}}) // ok
	eventMonitor.Trigger(&TEST_EVENT_ID_ANOTHER, t, 3, "gold")

	// emit
	eventMonitor.Emit()
	time.Sleep(10 * time.Millisecond)

	// release
	eventMonitor.Close()
}

func TestEventTimeListen(t *testing.T) {
	eventMonitor := NewEvent(&Option{
		InOrder:        false,
		TimeTickPeriod: time.Duration(10 * time.Millisecond),
		Submiter:       antsPool(), // 植入自己的携程池
	})

	// register
	eventMonitor.Register(EventHandle(handleLogin), &TEST_EVENT_ID_LOGIN)       // register fail
	eventMonitor.Register(EventHandle(handleLogout), &TEST_EVENT_ID_LOGOUT)     // register fail
	eventMonitor.Register(HandleConvert(anotherHandle), &TEST_EVENT_ID_ANOTHER) // register ok

	// trigger
	eventMonitor.Submit(&EventOper{&TEST_EVENT_ID_LOGIN, "", []any{t, 1, 2, "event"}}) // ok
	time.Sleep(20 * time.Millisecond)
	eventMonitor.Trigger(&TEST_EVENT_ID_ANOTHER, t, 3, "gold")

	// emit
	//eventMonitor.Emit()
	time.Sleep(100 * time.Millisecond)

	// release
	eventMonitor.Close()
}

// --------------------------testing used functions---------------------
var (
	TEST_EVENT_ID_LOGIN   EVENT_ID_STRING = "TEST.LOGIN"
	TEST_EVENT_ID_LOGOUT  EVENT_ID_STRING = "TEST.LOGOUT"
	TEST_EVENT_ID_ANOTHER EVENT_ID_INT    = 1001
)

func handleLogin(e *EventOper) {
	if len(e.Args) == 0 {
		panic(any("login handle can not recive the paramters"))
	}
	t, ok := e.Args[0].(*testing.T)
	if !ok {
		panic(any("please make your first pamar is a object of testing.T"))
	}
	t.Log("handleLogin was invocked")
}

func handleLogout(e *EventOper) {
	if len(e.Args) == 0 {
		panic(any("logout handle can not recive the paramters"))
	}
	t, ok := e.Args[0].(*testing.T)
	if !ok {
		panic(any("please make your first pamar is a object of testing.T"))
	}
	t.Log("handleLogout was invocked")
}

func handleTrace(e *EventOper) {
	if len(e.Args) == 0 {
		panic(any("trace handle can not recive the paramters"))
	}
	t, ok := e.Args[0].(*testing.T)
	if !ok {
		panic(any("please make your first pamar is a object of testing.T"))
	}
	t.Log("handleTrace was invocked")
}

func anotherHandle(args ...any) {
	if len(args) == 0 {
		panic(any("another handle can not recive the paramters"))
	}
	t, ok := args[0].(*testing.T)
	if !ok {
		panic(any("please make your first pamar is a object of testing.T"))
	}
	t.Log("anotherHandle<func(args ...any)> was invocked")
}
