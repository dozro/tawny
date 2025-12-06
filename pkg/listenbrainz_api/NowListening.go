package listenbrainz_api

import (
	"fmt"

	"github.com/dozro/tawny/internal/pkg/api_commons"
	"github.com/dozro/tawny/pkg/lfm_types"
	"github.com/dozro/tawny/pkg/listenbrainz_types"
	log "github.com/sirupsen/logrus"
)

func (User) GetCurrentTrackLfmCompat(username string) (*lfm_types.UserGetRecentTracks, error) {
	log.Debugf("getting recent tracks for %s ...", username)

	apiUrl := fmt.Sprintf("%s/1/user/%s/playing-now", baseUrl, username)

	data, err := api_commons.FetchJSON[listenbrainz_types.NowListeningWrapperOuter](apiUrl, "")

	if err != nil {
		return nil, err
	}

	var ret *lfm_types.UserGetRecentTracks
	ret, err = data.Payload.ConvertToLfm()
	return ret, err
}
