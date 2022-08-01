package ensure

const NoRaidList = ``

func EnsureNoRaidList() bool {
	return ensureExistence("noraidlist.txt", NoRaidList)
}
