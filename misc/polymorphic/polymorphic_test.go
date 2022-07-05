package polymorphic

import (
	"fmt"
	"testing"
)

// 多态方法
// 什么是多态? 根据类型的不同, 指代码可以根据类型的具体实现采取不同行为的能力
// 抽象=> 抽象行为=> 接口
// 类型=> 具体类型=> 结构体
//

type Notify interface {
	Notify()
}

type QQNotify struct {
	Name    string
	Message string
}

type MailNotify struct {
	Name    string
	Message string
}

type WechatNotify struct {
	Name    string
	Message string
}

func (t QQNotify) Notify() {
	fmt.Printf("%s-%s\n", t.Name, t.Message)
}

func (t MailNotify) Notify() {
	fmt.Printf("%s-%s\n", t.Name, t.Message)
}

func (t WechatNotify) Notify() {
	fmt.Printf("%s-%s\n", t.Name, t.Message)
}

func TestPolymorphic(t *testing.T) {
	qq := QQNotify{
		Name:    "QQ",
		Message: "我是QQ",
	}
	mail := MailNotify{
		Name:    "Email",
		Message: "我是邮箱",
	}

	wechat := WechatNotify{
		Name:    "Wechat",
		Message: "我是微信",
	}

	var notify Notify
	notify = qq
	notify.Notify()
	notify = mail
	notify.Notify()
	notify = wechat
	notify.Notify()
}
