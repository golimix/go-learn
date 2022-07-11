package single

// 1. 全局变量
var instanceHugy *Tool = new(Tool)

func GetInstanceHugy() *Tool {
	return instanceHugy
}

// 2. 包的初始化函数方式
func init() {
	instanceHugy = new(Tool)
}
