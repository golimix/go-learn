package typing

import "testing"

func TestTypingErrorSpec(t *testing.T) {
	println(`
		// 1. 预定义类型 error
		type error interface {
			Error() string
		}

		// 2. nil值表示没有错误

		// 3. 样例 
		func Read(f *File, b []byte) (n int, err error)
	`)
}

func TestTypingPanicError(t *testing.T) {
	println(`
		package runtime

		type Error interface {
			error
			// and perhaps other methods
		}
	`)
}
