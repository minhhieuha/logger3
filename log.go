package logger3

import (
	"fmt"
)

var Version string = "0.1.3"

func Log(mess string) {
	fmt.Println("[LOG] 3 " + mess)
}
