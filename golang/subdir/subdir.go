package subdir

import (
	"fmt"
	"runtime"
)

func TestSub() {
	pc, file, line, ok := runtime.Caller(0)
	fmt.Println(pc, file, line, ok)
}
