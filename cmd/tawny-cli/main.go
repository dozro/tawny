package main

import (
	"flag"
	"fmt"

	"github.com/dozro/tawny/pkg/lfm_api"
)

func main() {
	apiKey := flag.String("apikey", "", "api key")
	username := flag.String("username", "", "username")
	what := flag.String("what", "nowplaying", "what")
	flag.Parse()
	if *apiKey == "" || *username == "" || *what == "" {
		fmt.Println("Tawny cli tool, (C) 2025 itsrye.dev, Apache 2.0 Licensed, https://github.com/dozro/tawny")
		flag.Usage()
		return
	}
	if *what == "nowplaying" {
		np, _ := lfm_api.User{}.GetRecentTracks(lfm_api.UserGetArgsWithLimitPage{
			ApiKey:   *apiKey,
			UserName: *username,
			Limit:    1,
			Page:     -1,
		})
		if np.Track[0].NowPlaying {
			fmt.Printf("Now playing %s by %s\n", np.Track[0].Name, np.Track[0].Artist.Name)
		} else {
			fmt.Printf("Last playing %s by %s\n", np.Track[0].Name, np.Track[0].Artist.Name)
		}

	} else {
		fmt.Println("You need to specify what")
	}
}
