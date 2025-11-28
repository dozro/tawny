package main

import (
	"fmt"

	"github.com/dozro/tawny/pkg/tawny_sdk"
	log "github.com/sirupsen/logrus"
)

func main() {
	c := Flagread()
	log.SetLevel(log.DebugLevel)
	tawny := tawny_sdk.Tawny{}.NewTawny(tawny_sdk.TawnyCreationArgs{
		LastFMApiKey:  c.ApiKey,
		TawnyEndPoint: c.ApiEndpoint,
	})
	if c.Op == "user_tracks_current" {
		nl, err := tawny.GetNowListeningFor(c.Username)
		if err != nil {
			log.Fatal(err)
		}
		if nl.NowPlaying {
			fmt.Printf("%s is currently listening to \"%s\" by %s\n", c.Username, nl.Name, nl.Artist.Name)
		} else {
			fmt.Printf("%s was recently listening to \"%s\" by %s\n", c.Username, nl.Name, nl.Artist.Name)
		}
	}
}
