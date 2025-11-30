package server

import (
	"github.com/dozro/tawny/pkg/musicbrainz_api"
	"github.com/gin-gonic/gin"
)

func lookUpArtistByMbid(c *gin.Context) {
	artistMbid := c.Param("mbid")
	data, err := musicbrainz_api.ArtistLookupByMbid(artistMbid, true)
	if handleError(err, c) {
		return
	}
	render(c, 200, data)
}

func lookUpRecordingByMbid(c *gin.Context) {
	recordingMbid := c.Param("mbid")
	data, err := musicbrainz_api.RecordingLookupByMbid(recordingMbid, true)
	if handleError(err, c) {
		return
	}
	render(c, 200, data)
}
