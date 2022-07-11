package single

import "sync"

var dbLock sync.Mutex

// 双重检测模式
func GetInstance() *Tool {
	if instance == nil {
		dbLock.Lock()
		if instance == nil {
			instance = new(Tool)
		}
		dbLock.Unlock()
	}
	return instance
}
