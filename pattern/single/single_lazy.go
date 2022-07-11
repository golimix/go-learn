package single

import "sync"

type Tool struct {
	values int
}

var instance *Tool

// 懒汉式单例模式
func GetInstanceLazy() *Tool {
	if instance == nil {
		println("I am here")
		instance = &Tool{
			values: 100,
		}
	}
	return instance
}

// 懒汉线程安全式
var lock sync.Mutex

func GetInstanceLazySecure() *Tool {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = new(Tool)
	}
	return instance
}
