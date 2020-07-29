// Package internal provides common globally used functions.
package internal

import "news-service/config"

// IsSupportedLang checks if the specified language is in the config's SupportedLangs array
func IsSupportedLang(lang string) bool {
	for _, i := range config.SupportedLangs {
		if lang == i {
			return true
		}
	}

	return false
}
