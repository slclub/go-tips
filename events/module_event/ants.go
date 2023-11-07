package module_event

import (
	"errors"
	ants "github.com/panjf2000/ants/v2"
	"github.com/slclub/go-tips/logf"
	"runtime"
)

func antsPool() *ants.Pool {
	var ants_pool *ants.Pool
	if ants_pool == nil {
		var err error = errors.New("")
		ants_pool, err = ants.NewPool(runtime.NumCPU()*8, ants.WithLogger(logf.Log()))
		if err != nil {
			panic(any(err))
		}
	}
	return ants_pool
}
