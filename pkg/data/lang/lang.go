package lang

import (
	"fmt"
	"strings"

	ns "github.com/go-mclib/protocol/java_protocol/net_structures"
)

// Translate returns the English translation for a translation key, or empty string if not found.
func Translate(key string) string {
	return translations[key]
}

// TranslateComponent recursively flattens a TextComponent into an English
// string, resolving Translate keys using the built-in translation table.
func TranslateComponent(tc ns.TextComponent) string {
	var b strings.Builder
	writeTranslated(&b, &tc)
	return b.String()
}

func writeTranslated(b *strings.Builder, tc *ns.TextComponent) {
	if tc.Translate != "" {
		pattern := translations[tc.Translate]
		if pattern == "" {
			// no translation found, use the key
			b.WriteString(tc.Translate)
		} else {
			// resolve %s and %N$s placeholders with With arguments
			writeFormatted(b, pattern, tc.With)
		}
	} else {
		b.WriteString(tc.Text)
		b.WriteString(tc.Keybind)
		if tc.Score != nil {
			b.WriteString(tc.Score.Name)
		}
		b.WriteString(tc.Selector)
	}

	for _, child := range tc.Extra {
		writeTranslated(b, &child)
	}
}

// writeFormatted handles MC's Java-style format strings (%s, %1$s, %2$s, etc.)
func writeFormatted(b *strings.Builder, pattern string, args []ns.TextComponent) {
	seqIdx := 0 // sequential argument index for bare %s
	i := 0
	for i < len(pattern) {
		if pattern[i] != '%' || i+1 >= len(pattern) {
			b.WriteByte(pattern[i])
			i++
			continue
		}

		// found '%', look ahead
		j := i + 1

		// check for %% (literal percent)
		if pattern[j] == '%' {
			b.WriteByte('%')
			i = j + 1
			continue
		}

		// check for positional: %N$s
		if j+2 < len(pattern) && pattern[j] >= '1' && pattern[j] <= '9' && pattern[j+1] == '$' && pattern[j+2] == 's' {
			argIdx := int(pattern[j]-'0') - 1
			if argIdx >= 0 && argIdx < len(args) {
				writeTranslated(b, &args[argIdx])
			}
			i = j + 3
			continue
		}

		// bare %s
		if pattern[j] == 's' {
			if seqIdx < len(args) {
				writeTranslated(b, &args[seqIdx])
				seqIdx++
			}
			i = j + 1
			continue
		}

		// %d or other — just format the arg as string
		if pattern[j] == 'd' {
			if seqIdx < len(args) {
				b.WriteString(args[seqIdx].String())
				seqIdx++
			}
			i = j + 1
			continue
		}

		// unknown format specifier, output literally
		fmt.Fprintf(b, "%%%c", pattern[j])
		i = j + 1
	}
}
