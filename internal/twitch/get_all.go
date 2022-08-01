package twitch

import "github.com/nicklaw5/helix/v2"


func GetAllStreams(client *helix.Client, userLogins []string) []helix.Stream {
	var streams []helix.Stream
	var cursor string

	for cursor != "" {
		liveFromRaidlistRes, err := client.GetStreams(&helix.StreamsParams{UserLogins: userLogins, After: cursor})
		if err != nil {
			panic(err)
		}
		cursor = liveFromRaidlistRes.Data.Pagination.Cursor
		streams = append(streams, liveFromRaidlistRes.Data.Streams...)
	}

	return streams
}

func GetAllFollowedStreams(client *helix.Client, streamerId string) []helix.Stream {
	var streams []helix.Stream
	var cursor string

	for cursor != "" {
		liveFromRaidlistRes, err := client.GetFollowedStream(&helix.FollowedStreamsParams{UserID: streamerId, After: cursor})
		if err != nil {
			panic(err)
		}
		cursor = liveFromRaidlistRes.Data.Pagination.Cursor
		streams = append(streams, liveFromRaidlistRes.Data.Streams...)
	}

	return streams
}
