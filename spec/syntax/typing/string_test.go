package typing

import (
	"fmt"
	"strings"
	"testing"
)

func TestTypingString(t *testing.T) {
	var err interface{}
	defer func() {
		if err = recover(); err != nil {
			fmt.Println("main ==》 ", err)
		}
	}()
	a := "abcde"
	fmt.Printf("%c\t%c\t%c\t%c\t%c\t\n", a[0], a[1], a[2], a[3], a[4])
	fmt.Printf("%c\t%c\t%c\t%c\t%c\t%c\n", a[0], a[1], a[2], a[3], a[4], a[5])
}

// Values of string type are immutable
func changeString(email string) string {
	fmt.Printf("email: %v\n", email)
	email = "normal"
	return email
}

func changeStringPoint(email *string) string {
	*email = "point"
	return *email
}

// 测试字符串参数传递
func TestTypingStringParams(t *testing.T) {
	fmt.Println("type string 的值是immutable")
	user := "init"
	user1 := changeString(user)
	fmt.Printf("%p,%p\n", &user, &user1)
	user2 := changeStringPoint(&user)
	fmt.Printf("%p,%p,%p\n", &user, &user1, &user2)

	user3 := ([]byte)("init")
	println(&user3[0], "\n")
}

func TestTypingStringCompare(t *testing.T) {
	a := "A"
	b := "a"
	if strings.Compare(strings.ToLower(a), strings.ToLower(b)) == 0 {
		println("a = b ignore caption")
	}
}
