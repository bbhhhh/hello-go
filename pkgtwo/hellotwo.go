package pkgtwo

import (
	"fmt"
	"os"
)

const A = "a"

func Print2(a ...interface{}) (n int) {
	fmt.Fprintln(os.Stdout, a...)
	return 0
}
