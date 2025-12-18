package common_types

type MetaInformation struct {
	CachingInformation CachingInformation `json:"caching_information" xml:"caching-information"`
	CompatInformation  CompatInformation  `json:"compat_information" xml:"compat-information"`
}
