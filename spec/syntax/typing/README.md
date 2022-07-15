## 类型分类
可参考 https://www.jianshu.com/p/ce307b8e9772

### 命名类型
    int, int64, float32, string, bool等。这些已经是GO中预先声明好的类型
    通过类型声明(type declaration)创建的所有类型都是命名类型

结论: 一个命名类型一定与其他类型不同

```go
var i int // named type
type myInt int // named type
var b bool // named type
```

### 未命名类型
虽然他们没有名字，但却有一个类型字面量(type literal)来描述他们由什么构成

    数组，结构体，指针，函数，接口，切片，map，通道 都是未命名类型

```go
[]string // unnamed type
map[string]string // unnamed type
[10]int // unnamed type
```

### 基础类型
任何类型T都有基本类型

如果T 是预先声明类型：boolean, numeric, or string（布尔，数值，字符串）中的一个，或者是一个类型字面量(type literal)，他们对应的基础类型就是T自身
否则，T的基础类型就是 T所引用的那个类型的类型声明(type declaration)


### 类型转换
类型转换, 如果他们基础类型的字面量在结构上是等价的,他们就是一致的类型

```go
package main

type Meter struct {
    value int64
}

type Centimeter struct {
    value int32
}

func main() {
    cm := Centimeter{
        value: 1000,
    }

    var m Meter
    m = Meter(cm)
    print(m.value)
    cm = Centimeter(m)
    print(cm.value)
}
```
