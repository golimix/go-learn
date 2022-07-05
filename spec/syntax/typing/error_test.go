package typing

import "testing"

/* 函数调用 => 发生错误 => 错误表征 => 错误返回 => 错误处理
 函数包: import errors
 函数调用: 不一定会成功,必定存在发生错误场景
 错误表征:
    Java: Exception
    是否发生: 类型 error (空:成功, 非空:失败)
    详细信息: fmt.Println(err) 或者 fmt.Println("%v", err)
 错误返回: 多返回值
 错误案例: 查询缓存可能会失败
 错误处理:
    传统: 处理控制流 throw catch
    Go语言: 正常控制流 panic recover
 处理策略:
    方式一、直接返回错误 -> 子例程错误变为主调例程的错误
    resp, err := http.Get(url)
    if err != nil {
 	     return nil, err
    }
    建新的错误消息
    doc, err := html.Parse(resp.Body)
    resp.Body.Close()
    if err != nil {
	   // 错误处理链条
       return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
    }
	TODO 错误及其处理
	方式二、重试
	方式三、优雅的退出
	方法四、只记录下错误信息,然后程序继续运行
*/
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
