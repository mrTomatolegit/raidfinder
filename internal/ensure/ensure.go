package ensure

import (
	"os"
)

func ensureExistence(filename string, filltext string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file.Write([]byte(filltext))
		return false
	}
	return true
}
