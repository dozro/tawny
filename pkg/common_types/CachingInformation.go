package common_types

type CachingInformation struct {
	CacheHit     bool   `json:"cache_hit" xml:"cache-hit"`
	CacheBackend string `json:"cache_backend", xml:"cache-backend"`
}
