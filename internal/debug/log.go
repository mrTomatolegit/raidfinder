package debug

import "fmt"

func Log(a ...any) {
	if Enabled {
		fmt.Println(a...)
	}
}
