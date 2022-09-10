package main

import (
	"fmt"

	"github.com/mrTomatolegit/raid-finder/internal/ensure"
	"github.com/mrTomatolegit/raid-finder/internal/twitch"
	"github.com/mrTomatolegit/raid-finder/internal/util"
	"github.com/mrTomatolegit/raid-finder/internal/debug"
	"github.com/nicklaw5/helix/v2"
)

func ensureFiles() bool {
	raidlistExisted := ensure.EnsureRaidList()
	noraidlistExisted := ensure.EnsureNoRaidList()
	return raidlistExisted && noraidlistExisted
}

func main() {
	if !ensureFiles() {
		fmt.Println("Created raidlist.txt and/or noraidlist.txt")
	}

	client := util.LoadClient()

	streamerResp, err := client.GetUsers(&helix.UsersParams{})
	if err != nil {
		panic(err)
	}
	streamer := streamerResp.Data.Users[0]

	raidlist := util.LoadList("raidlist.txt")
	noraidlist := util.LoadList("noraidlist.txt")
	raidlist = util.Filter(&raidlist, func(e string, _ int) bool {
		return !util.Contains(&noraidlist, e)
	})

	fmt.Println("Looking for live streams...")

	liveFromRaidListChan := make(chan []helix.Stream)
	go twitch.GetAllStreamsAsync(client, raidlist, liveFromRaidListChan)
	liveFromFollowingChan := make(chan []helix.Stream)
	go twitch.GetAllFollowStreamsAsync(client, streamer.ID, noraidlist, liveFromFollowingChan)
	liveFromSearchChan := make(chan []helix.Stream)
	go twitch.GetSameGameStreamsAsync(client, streamer.ID, noraidlist, liveFromSearchChan)

	liveFromRaidlist := <-liveFromRaidListChan
	liveFromFollowing := <-liveFromFollowingChan
	liveFromSearch := <-liveFromSearchChan

	if len(liveFromRaidlist) > 0 {
		fmt.Println("\nLive from raidlist:")
		for _, stream := range liveFromRaidlist {
			fmt.Printf(util.StringifyStream(stream) + "\n")
		}
	} else {
		if len(raidlist) == 0 {
			fmt.Print("\nRaidlist is empty\n")
		} else {
			fmt.Print("\nNo live streamers from raidlist\n")
		}
	}

	if len(liveFromFollowing) > 0 {
		fmt.Println("\nLive from following:")
		for _, stream := range liveFromFollowing {
			fmt.Printf(util.StringifyStream(stream) + "\n")
		}
	} else {
		fmt.Print("\nNo live streamers from following\n")
	}

	if len(liveFromSearch) > 0 {
		fmt.Println("\nLive from following:")
		for _, stream := range liveFromFollowing {
			fmt.Printf(util.StringifyStream(stream) + "\n")
		}
	} else {
		fmt.Print("\nNo other live streamers that are playing the same game\n")
	}

	if !debug.Enabled {
		util.WaitForKeypress()
	}
}
