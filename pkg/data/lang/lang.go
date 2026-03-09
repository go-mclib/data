package lang

import ns "github.com/go-mclib/protocol/java_protocol/net_structures"

func init() {
	ns.TranslateFunc = Translate
}

// Translate returns the English translation for a translation key, or empty string if not found.
func Translate(key string) string {
	return translations[key]
}
