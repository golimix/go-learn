package stdlib_os

import (
	"fmt"
	"os"
	"testing"
)

func TestXxx(t *testing.T) {
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}
