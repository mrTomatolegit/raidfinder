package twitch

import (
	"github.com/mrTomatolegit/raid-finder/internal/debug"
	"github.com/nicklaw5/helix/v2"
)

func GetAllStreams(client *helix.Client, userLogins []string) []helix.Stream {
	debug.Log("Fetching streams for users", userLogins)
	var streams []helix.Stream
	var cursor string

	liveFromRaidlistRes, err := client.GetStreams(&helix.StreamsParams{UserLogins: userLogins})
	if err != nil {
		panic(err)
	}
	cursor = liveFromRaidlistRes.Data.Pagination.Cursor
	streams = append(streams, liveFromRaidlistRes.Data.Streams...)
	debug.Log("Fetching streams for users result 0", streams)

	for cursor != "" {
		liveFromRaidlistRes, err := client.GetStreams(&helix.StreamsParams{UserLogins: userLogins, After: cursor})
		if err != nil {
			panic(err)
		}
		cursor = liveFromRaidlistRes.Data.Pagination.Cursor
		streams = append(streams, liveFromRaidlistRes.Data.Streams...)
	}
	debug.Log("Fetching streams for users result final", streams)

	return streams
}

func GetAllFollowedStreams(client *helix.Client, streamerId string) []helix.Stream {
	debug.Log("Fetching followed streams for", streamerId)
	var streams []helix.Stream
	var cursor string

	liveFromRaidlistRes, err := client.GetFollowedStream(&helix.FollowedStreamsParams{UserID: streamerId})
	if err != nil {
		panic(err)
	}
	cursor = liveFromRaidlistRes.Data.Pagination.Cursor
	streams = append(streams, liveFromRaidlistRes.Data.Streams...)

	debug.Log("Fetching followed streams for result 0", streams)

	for cursor != "" {
		liveFromRaidlistRes, err := client.GetFollowedStream(&helix.FollowedStreamsParams{UserID: streamerId, After: cursor})
		if err != nil {
			panic(err)
		}
		cursor = liveFromRaidlistRes.Data.Pagination.Cursor
		streams = append(streams, liveFromRaidlistRes.Data.Streams...)
	}

	debug.Log("Fetching followed streams for result final", streams)

	return streams
}
