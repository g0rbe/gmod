package validator

import "strings"

func emailLocal(v string) bool {

	lenV := len(v)

	if lenV == 0 || lenV > 64 {
		return false
	}

	last := byte('.')

	for i := 0; i < lenV; i++ {

		switch {
		case v[i] >= 'a' && v[i] <= 'z':
			last = v[i]
		case v[i] >= 'A' && v[i] <= 'Z':
			last = v[i]
		case v[i] >= '0' && v[i] <= '9':
			last = v[i]
		case v[i] == '#' || v[i] == '$' || v[i] == '%' || v[i] == '&' || v[i] == '\'' || v[i] == '*':
			last = v[i]
		case v[i] == '+' || v[i] == '-' || v[i] == '/' || v[i] == '=' || v[i] == '?' || v[i] == '^':
			last = v[i]
		case v[i] == '_' || v[i] == '`' || v[i] == '{' || v[i] == '|' || v[i] == '}' || v[i] == '~':
			last = v[i]
		case v[i] == '.':
			// local part cant start, end and contains two dot or more dots in a row.
			if last == '.' {
				return false
			}
			last = v[i]

		default:
			return false
		}
	}

	return last != '.'
}

func Email(v string) bool {

	// 64 is Unicode '@'
	sep := strings.IndexRune(v, 64)
	if sep == -1 {
		return false
	}

	// Check local part
	if !emailLocal(v[:sep]) {
		return false
	}

	// Empty domain (eg.: user@)
	if len(v)-1 == sep {
		return false
	}

	// Invaid domain
	if !Domain(v[sep+1:]) {
		return false
	}

	return true
}
