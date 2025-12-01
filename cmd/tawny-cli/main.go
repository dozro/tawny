package main

import (
	"fmt"

	"github.com/dozro/tawny/pkg/tawny_sdk"
	log "github.com/sirupsen/logrus"
)

func main() {
	c := Flagread()
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	if c.Debug {
		log.SetLevel(log.DebugLevel)
		log.Debug("Debug logging enabled")
	}
	tawny := tawny_sdk.Tawny{}.NewTawny(tawny_sdk.TawnyCreationArgs{
		LastFMApiKey:  c.ApiKey,
		TawnyEndPoint: c.ApiEndpoint,
		HMACSecretKey: c.HMACSecret,
	})
	log.Debug(tawny)
	if c.Op == "current" {
		log.Debugf("getting now listening for %s", c.Username)
		nl, err := tawny.SecureNowListeningFor(c.Username)
		if err != nil {
			log.Fatal(err)
		}
		if nl.NowPlaying {
			fmt.Printf("%s is currently listening to \"%s\" by %s\n", c.Username, nl.Name, nl.Artist.Name)
		} else {
			fmt.Printf("%s was recently listening to \"%s\" by %s\n", c.Username, nl.Name, nl.Artist.Name)
		}
	} else if c.Op == "userinfo" {

	}
}
