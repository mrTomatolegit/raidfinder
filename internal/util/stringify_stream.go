package util

import (
	"fmt"
	"github.com/nicklaw5/helix/v2"
	"math"
	"time"
)

func StringifyStream(stream helix.Stream) string {
	var userName = stream.UserName
	var gameName = stream.GameName
	var since = time.Since(stream.StartedAt)
	var sinceHours = int(math.Floor(since.Hours()))
	var sinceMinutes = int(math.Floor(since.Minutes())) - sinceHours*60
	var viewerCount = stream.ViewerCount
	return fmt.Sprintf("- %s [%dh %dm] (%d viewers) is playing %s", userName, sinceHours, sinceMinutes, viewerCount, gameName)
}
