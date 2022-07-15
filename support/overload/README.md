### 说明
Go语言默认是不支持重载机制,例如:下面的例子,在同一个作用域范围内,只能有一个方法

### What
在同一作用域下面,Handler方法不能重复出现

```go
func Handler()  {   
}

func Handler(timeOut time.Duration) { 
}

func Handler(timeOut time.Duration, retry int) {
}
```

### Why


### How
通过变长参数 + 链式处理, 传递参数指针,返回的还是参数指针


```go
func Handler(op ...interface{}) {

}
```