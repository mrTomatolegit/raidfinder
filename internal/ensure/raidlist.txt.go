package ensure

const RaidList = ``

func EnsureRaidList() bool {
	return ensureExistence("raidlist.txt", RaidList)
}
