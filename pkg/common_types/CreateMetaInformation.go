package common_types

const COMPAT_DATE = "2025-12-18"

func CreateMetaInformation(cachehit bool, cacheBackend string) MetaInformation {
	return MetaInformation{
		CompatInformation: CompatInformation{
			CompatDate:    COMPAT_DATE,
			CompatVersion: "",
		},
		CachingInformation: CachingInformation{
			CacheHit:     cachehit,
			CacheBackend: cacheBackend,
		},
	}
}
