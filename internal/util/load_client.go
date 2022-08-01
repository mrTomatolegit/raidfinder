package util

import (
	"fmt"
	"github.com/mrTomatolegit/raid-finder/internal/twitch/oauth2"
	"github.com/nicklaw5/helix/v2"
)

func applyCachedToken(client *helix.Client) bool {
	token, err := oauth2.LoadCachedToken()
	if err != nil {
		return false;
	}

	isValid, _, err := client.ValidateToken(token)
	if err != nil {
		return false
	}

	if (isValid) {
		client.SetUserAccessToken(token)
		return true
	} else {
		fmt.Println("Twitch token is invalid, please re-authorize")
		return false
	}
}

func LoadClient() *helix.Client {
	client, err := helix.NewClient(&helix.Options{
		ClientID:    "aq6y0emkp6xcfwa5bn2ier93yxarw7",
		RedirectURI: "http://localhost:42069",
	})
	if err != nil {
		panic(err)
	}

	cached := applyCachedToken(client)

	if !cached {
		done := oauth2.AwaitUserAccessToken(client)
		<-done
		oauth2.SaveCachedToken(client.GetUserAccessToken())
	}

	return client
}
