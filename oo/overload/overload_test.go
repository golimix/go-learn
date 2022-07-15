package overload

import (
	"testing"
	"time"
)

// 接收Options指针的函数
type OptionsFunc func(*Options)

// 选项参数
type Options struct {
	TimeOut     time.Duration
	RetryMaxNum int
}

// 核心-链式处理-类似于Builder模式
// 链式的Options指针函数
func loadOp(option ...OptionsFunc) *Options {
	options := new(Options)
	for _, e := range option {
		e(options)
	}
	return options
}

func Handler(option ...OptionsFunc) {
	op := loadOp(option...)
	println(op.TimeOut)
	println(op.RetryMaxNum)
}

func TestOverload(t *testing.T) {
	Handler()
	Handler(func(options *Options) {
		options.TimeOut = time.Millisecond
	})

	Handler(func(options *Options) {
		options.RetryMaxNum = 2
	})

	Handler(func(options *Options) {
		options.RetryMaxNum = 3
	}, func(options *Options) {
		options.TimeOut = 3 * time.Millisecond
	})
}
