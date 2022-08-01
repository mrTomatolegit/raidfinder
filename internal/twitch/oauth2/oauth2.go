package oauth2

import (
	"embed"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/nicklaw5/helix/v2"
)

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		panic(err)
	}

}

func LoadCachedToken() (string, error) {
	_, err := os.Stat("cached.token")
	if err != nil {
		return "", errors.New("cached.token not found")
	}
	bytes, err := os.ReadFile("cached.token")
	if err != nil {
		panic(err)
	}
	token := string(bytes)

	if token == "" {
		return "", errors.New("cached.token is empty")
	}

	return token, nil
}

func SaveCachedToken(token string) {
	file, err := os.Create("cached.token")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte(token))
	fmt.Println("Twitch token saved to cached.token")
}

//go:embed html/*
var f embed.FS

func AwaitUserAccessToken(client *helix.Client) chan struct{} {
	fmt.Println("Sending you to the twitch login page")
	done := make(chan struct{})

	server := http.Server{Addr: ":42069", Handler: nil}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Has("access_token") {
			token := r.URL.Query().Get("access_token")
			client.SetUserAccessToken(token)
			w.Header().Add("Content-Type", "text/html")
			bts, err := f.ReadFile("html/finish.html")
			if err != nil {
				panic(err)
			}
			w.Write([]byte(string(bts)))
			// w.Write([]byte(`<body></body><script>setTimeout(window.close())</script>`))
			done <- struct{}{}
			close(done)
		} else if !r.URL.Query().Has("error") {
			w.Header().Add("Content-Type", "text/html")

			bts, err := f.ReadFile("html/hashtransform.html")
			if err != nil {
				panic(err)
			}
			w.Write([]byte(string(bts)))
			// w.Write([]byte(`<script>document.location.search = '?' + document.location.hash.substring(1);document.location.hash = undefined;</script>`))
		} else {
			w.Write([]byte("An error has occured!"))
			fmt.Println("Error", r.URL.Query().Get("error_description"))
			os.Exit(1)
		}
	})

	go server.ListenAndServe()

	url := client.GetAuthorizationURL(&helix.AuthorizationURLParams{
		ResponseType: "token",
		Scopes:       []string{"user:read:follows"},
		State:        "state",
		ForceVerify:  false,
	})

	openbrowser(url)

	return done
}
