package literal

import (
	"fmt"
	"testing"
)

func TestInteger(t *testing.T) {
	println("包括二进制,十进制,八进制,十六进制")
	PrintInteger(0b100100)
	PrintInteger(36)
	PrintInteger(0o44)
	PrintInteger(0x24)
}

func PrintInteger(a int) {
	fmt.Printf("二进制=%#b\t十进制=%d\t八进制=%#o\t十六进制=%#x\n", a, a, a, a)
}
