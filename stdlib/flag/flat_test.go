package stdlib_flag

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	println(`
		Package flag implements command-line flag parsing. 
		实现命令行参数解析	
	`)
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
	// 定义flag
	name := flag.String("name", "张三", "姓名")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "婚否")
	delay := flag.Duration("d", 0, "时间间隔")
	fmt.Printf("%s-%d-%t-%s\n", *name, *age, *married, *delay)
}

func TestFlagTypeVar(t *testing.T) {
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "张三", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "时间间隔")
	fmt.Printf("%s-%d-%t-%s\n", name, age, married, delay)
}
