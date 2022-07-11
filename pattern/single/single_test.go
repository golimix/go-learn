// 单例模式
package single

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestPatternSingleLazy(t *testing.T) {
	for i := 0; i < 10; i++ {
		go fmt.Printf("%3d GetInstanceHugy(): %p\n", runtime.NumGoroutine(), GetInstanceLazy())
	}
	// 线程等待
	time.Sleep(time.Second * 10)
}

func TestPatternSingleLazySecure(t *testing.T) {
	for i := 0; i < 10; i++ {
		go fmt.Printf("%3d GetInstanceHugy(): %p\n", runtime.NumGoroutine(), GetInstanceLazySecure())
	}
	// 线程等待
	time.Sleep(time.Second * 10)
}
