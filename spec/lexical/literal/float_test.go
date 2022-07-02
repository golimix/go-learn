package literal

import (
	"fmt"
	"testing"
)

func TestFloat(t *testing.T) {
	// 标准十进制
	fmt.Printf("%f\n", 72.40)
	// 二进制
	fmt.Printf("%f\n", 0x1p-2) // =1/2*2
	fmt.Printf("%f\n", 0x2p10) // = 2*2的10次方=2048
	// 十进制
	fmt.Printf("%f\n", 1.e+0)
	fmt.Printf("%f\n", 1.e+0)
}
