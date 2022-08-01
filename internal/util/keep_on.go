package util

import (
	"fmt"
	"bufio"
	"os"
)

func WaitForKeypress() {
	fmt.Print("\nPress the ENTER key to close...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
