package server

import (
	"github.com/dozro/tawny/pkg/musicbrainz_api"
	"github.com/gin-gonic/gin"
)

func lookUpArtistByMbid(c *gin.Context) {
	artistMbid := c.Param("artistMbid")
	data, err := musicbrainz_api.ArtistLookupByMbid(artistMbid, true)
	if handleError(err, c) {
		return
	}
	render(c, 200, data)
}
