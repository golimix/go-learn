package literal

import (
	"fmt"
	"testing"
)

// 符文字面量字母
func TestLiteralRuneLetter(t *testing.T) {
	// 输出letter
	Print('a')
	Print('_')
}

// 只有以下的转义字符是有效的,其他的都会被认为是非法输入
func TestLiteralRuneEscape(t *testing.T) {
	Print('\a')
	Print('\b')
	Print('\f')
	Print('\n')
	Print('\r')
	Print('\t')
	Print('\v')
	Print('\\')
	Print('\'')
}

// byte_value = octal_byte_value | hex_byte_value .
func TestLiteralRuneByte(t *testing.T) {
	// 八进制 octal_byte_value = `\` octal_digit octal_digit octal_digit .
	Print('\000')
	// 十六进制 hex_byte_value   = `\` "x" hex_digit hex_digit .
	Print('\xff')
}

func TestLiteralRuneLitteUnicode(t *testing.T) {
	Print('\u12e4')
	// Print('\uDFFF') 非法范围
}

func TestLiteralRuneBigUnicode(t *testing.T) {
	Print('\U00101234')
	// Print('\U00110000') 非法范围
}

func Print(a rune) {
	fmt.Println(a)
	fmt.Printf("%c\n", a)
}
