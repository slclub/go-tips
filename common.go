package tips

import (
	"fmt"
	"github.com/slclub/go-tips/convtype"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var Rand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// 获取routine ID
func GetGoroutineId() int {
	defer func() {
		if err := recover(); err != any(nil) {
			log.Print("panic recover:panic info:", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(any(fmt.Sprintf("cannot get goroutine id: %v", err)))
	}
	return id
}

// 获取可执行文件的绝对根路径
func GetRootPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Print(err)
	}
	// For testing command.
	if dir[:4] == "/tmp" {
		dir, err = os.Getwd()
	}
	return dir
}

// force convert success
// It will be zere if any error happend
func Any2Int64(v any) int64 {
	return convtype.Any2Int64(v)
}

func Int(v any) int {
	return int(Any2Int64(v))
}
