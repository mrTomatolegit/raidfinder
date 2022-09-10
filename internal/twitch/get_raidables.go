package twitch

import (
	"github.com/mrTomatolegit/raid-finder/internal/debug"
	"github.com/mrTomatolegit/raid-finder/internal/util"
	"github.com/nicklaw5/helix/v2"
)

func GetAllStreamsAsync(client *helix.Client, userlogins []string, c chan []helix.Stream) {
	defer close(c)
	var streams []helix.Stream
	if len(userlogins) > 0 {
		streams = GetAllStreams(client, userlogins)
	}
	c <- streams
}

func GetAllFollowStreamsAsync(client *helix.Client, streamerId string, noraidlist []string, c chan []helix.Stream) {
	defer close(c)
	streams := GetAllFollowedStreams(client, streamerId)
	validStreams := util.Filter(&streams, func(stream helix.Stream, _ int) bool {
		return !util.Contains(&noraidlist, stream.UserLogin)
	})
	c <- validStreams
}

func GetSameGameStreamsAsync(client *helix.Client, streamerId string, noraidlist []string, c chan []helix.Stream) {
	defer close(c)
	debug.Log("Finding streams of the same game as", streamerId)
	streamerStreamResp, err := client.GetStreams(&helix.StreamsParams{UserIDs: []string{streamerId}})
	if err != nil {
		panic(err)
	}
	if len(streamerStreamResp.Data.Streams) == 0 {
		c <- []helix.Stream{}
		return
	}
	streamerStream := streamerStreamResp.Data.Streams[0]
	gameId := streamerStream.GameID
	language := streamerStream.Language
	otherStreamsResp, err := client.GetStreams(&helix.StreamsParams{GameIDs: []string{gameId}, Language: []string{language}})

	if err != nil {
		panic(err)
	}
	otherStreams := otherStreamsResp.Data.Streams
	otherStreams = util.Filter(&otherStreams, func(stream helix.Stream, _ int) bool {
		return !util.Contains(&noraidlist, stream.UserLogin)
	})
	debug.Log("Finding streams of the same game result final", otherStreams)
	c <- otherStreams
}
