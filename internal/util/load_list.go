package util

import (
	"os"
	"strings"
)

func LoadList(path string) []string {
	raidlistFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	raidlistArr := strings.Split(string(raidlistFile), "\n")

	for i := 0; i < len(raidlistArr); i++ {
		trimmed := strings.Trim(raidlistArr[i], " ")
		if len(trimmed) == 0 {
			Splice(&raidlistArr, i, 1)
			i--
		} else {
			raidlistArr[i] = strings.ToLower(trimmed)
		}
	}

	return raidlistArr
}
