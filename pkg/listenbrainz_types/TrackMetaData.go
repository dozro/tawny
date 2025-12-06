package listenbrainz_types

type TrackMetaData struct {
	AdditionalInfo TrackMetaDataAdditionalInfo `json:"additional_info"`
	ArtistName     string                      `json:"artist_name"`
	ReleaseName    string                      `json:"release_name"`
	TrackName      string                      `json:"track_name"`
	RecordingMbid  string                      `json:"recording_mbid"`
	TrackMbid      string                      `json:"track_mbid"`
	ArtistMbids    []string                    `json:"artist_mbids"`
	WorkMbids      []string                    `json:"work_mbids"`
	SpotifyId      string                      `json:"spotify_id"`
}
type TrackMetaDataAdditionalInfo struct {
	DurationMS              int    `json:"duration_ms"`
	SubmissionClient        string `json:"submission_client"`
	SubmissionClientVersion string `json:"submission_client_version"`
}
