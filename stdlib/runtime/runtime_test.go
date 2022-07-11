package stdlib_runtime

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGoRoutineNum(t *testing.T) {
	// 打印当前的线程
	fmt.Printf("%d", runtime.NumGoroutine())
}
