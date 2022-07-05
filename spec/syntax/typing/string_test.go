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

// 字符串比较
// 推荐使用 strings.EqualFold方法进行字符串比较
func TestTypingStringCompare(t *testing.T) {
	// 1. 区分大小写-直接使用 ==
	srcString := "This a string"
	destString := "this a string"
	if srcString == destString {
		fmt.Println("Equals")
	} else {
		fmt.Println("Not Equals")
	}

	// 2. 借助toLower已废弃
	// if strings.ToLower(srcString) == strings.ToLower(destString) {
	// 	fmt.Println("Equals")
	// } else {
	// 	fmt.Println("Not Equals")
	// }

	if strings.Compare(strings.ToLower(srcString), strings.ToLower(destString)) == 0 {
		fmt.Println("Equals")
	} else {
		fmt.Println("Not Equals")
	}

	if strings.EqualFold(srcString, destString) == true {
		fmt.Println("Equals")
	} else {
		fmt.Println("Not Equals")
	}
}
