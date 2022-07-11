package single

import "sync"

var onceInstance *Tool
var once sync.Once

// 本质上也是双重检测,但是检查更间接
func GetInstanceOnce() *Tool {
	once.Do(func() {
		onceInstance = new(Tool)
	})
	return onceInstance
}
